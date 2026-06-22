package library

import (
	"os"
	"path/filepath"
	"strings"

	"tunedock/internal/model"
)

func ReadAudioMetadata(audioPath string) (*model.Track, error) {
	info, err := os.Stat(audioPath)
	if err != nil {
		return nil, err
	}

	codec := codecFromExt(audioPath)

	return &model.Track{
		Title:      titleFromFileName(audioPath),
		FilePath:   audioPath,
		FileSize:   info.Size(),
		Codec:      codec,
		IsLossless: false,
		IsHiRes:    false,
	}, nil
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
