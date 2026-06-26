package library

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dhowden/tag"
	"github.com/google/uuid"
	"go.senan.xyz/taglib"

	"tunedock/internal/model"
	"tunedock/internal/repository"
)

func ReadAudioMetadata(database *sql.DB, audioPath string) (*model.Track, error) {
	info, err := os.Stat(audioPath)
	if err != nil {
		return nil, err
	}

	codec := codecFromExt(audioPath)

	file, err := os.Open(audioPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// タグが読めなかった場合でも最低限登録できるようにする
	trackTitle := titleFromFileName(audioPath)
	tags, taglibErr := taglib.ReadTags(audioPath)
	if taglibErr != nil {
		tags = map[string][]string{}
	}
	discNumber := 1

	var albumID *string
	var trackNumber *int
	var durationMS *int
	var sampleRate *int
	var bitrate *int

	m, err := tag.ReadFrom(file)
	if err == nil {
		if title := strings.TrimSpace(m.Title()); title != "" {
			trackTitle = title
		}

		trackNumber = getTrackNumber(m, tags)

		if rawDiscNumber := rawString(m.Raw(), "DISCNUMBER"); rawDiscNumber != "" {
			if parsedDiscNumber := parseNumber(rawDiscNumber); parsedDiscNumber != nil {
				discNumber = *parsedDiscNumber
			}
		}

		albumTitle := strings.TrimSpace(m.Album())
		if albumTitle != "" {
			album, err := repository.GetAlbumByName(database, albumTitle)

			if err != nil {
				fmt.Printf("[ERROR]: %s", err)
				if errors.Is(err, sql.ErrNoRows) {
					albumYear := m.Year()
					var albumYearPtr *int
					if albumYear > 0 {
						albumYearPtr = &albumYear
					}

					newAlbum := &model.Album{
						ID:      uuid.NewString(),
						Title:   albumTitle,
						Year:    albumYearPtr,
						CoverID: uuid.NewString(),
					}
					
					fmt.Printf("\n--- Create New Album DB ---\n%v", newAlbum)
					if err := repository.UpsertAlbum(database, newAlbum); err != nil {
						return nil, err
					}

					album = newAlbum
				} else {
					return nil, err
				}
			}

			if album != nil {
				albumID = &album.ID
			}
		}
	}

	props, err := taglib.ReadProperties(audioPath)
	if err == nil {
		if props.Length > 0 {
			value := int(props.Length.Milliseconds())
			durationMS = &value
		}

		if props.SampleRate > 0 {
			value := int(props.SampleRate)
			sampleRate = &value
		}

		if props.Bitrate > 0 {
			value := int(props.Bitrate)
			bitrate = &value
		}
	}

	

	return &model.Track{
		AlbumID:     albumID,
		Title:       trackTitle,
		DiscNumber:  discNumber,
		TrackNumber: trackNumber,
		DurationMS:  durationMS,

		FilePath: audioPath,
		FileSize: info.Size(),

		Codec:      codec,
		SampleRate: sampleRate,
		Bitrate:    bitrate,

		IsLossless: isLosslessCodec(codec),
		IsHiRes:    isHiRes(sampleRate, nil),
	}, nil
}

func getTrackNumber(m tag.Metadata, tags map[string][]string) *int {
	if m != nil {
		if n, _ := m.Track(); n > 0 {
			return &n
		}
	}

	if s := firstTag(tags, taglib.TrackNumber); s != "" {
		return parseNumber(s)
	}

	if m != nil {
		if s := rawString(m.Raw(), "TRACKNUMBER"); s != "" {
			return parseNumber(s)
		}
		if s := rawString(m.Raw(), "TRACK"); s != "" {
			return parseNumber(s)
		}
		if s := rawString(m.Raw(), "TRCK"); s != "" {
			return parseNumber(s)
		}
	}

	return nil
}

func firstTag(tags map[string][]string, key string) string {
	values := tags[key]
	if len(values) == 0 {
		return ""
	}

	return strings.TrimSpace(values[0])
}

func rawString(raw map[string]interface{}, key string) string {
	value, ok := raw[key]
	if !ok {
		return ""
	}

	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v)
	case []string:
		if len(v) == 0 {
			return ""
		}
		return strings.TrimSpace(v[0])
	default:
		return ""
	}
}

func parseNumber(s string) *int {
	if s == "" {
		return nil
	}

	// "3/12" のような形式に対応
	s = strings.Split(s, "/")[0]
	s = strings.TrimSpace(s)

	n, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}

	return &n
}

func isLosslessCodec(codec string) bool {
	switch strings.ToLower(codec) {
	case "flac", "wav", "aiff", "aif", "alac":
		return true
	default:
		return false
	}
}

func isHiRes(sampleRate *int, bitDepth *int) bool {
	if sampleRate != nil && *sampleRate > 48000 {
		return true
	}

	if bitDepth != nil && *bitDepth >= 24 {
		return true
	}

	return false
}

func titleFromFileName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)

	return strings.TrimSuffix(base, ext)
}

func codecFromExt(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	return strings.TrimPrefix(ext, ".")
}
