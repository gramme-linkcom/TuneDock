PRAGMA foreign_keys = ON;

INSERT OR IGNORE INTO artists (
  id,
  name,
  sort_name
) VALUES
(
  '11111111-1111-1111-1111-111111111111',
  'Sample Artist',
  'Sample Artist'
);

INSERT OR IGNORE INTO albums (
  id,
  title,
  year,
  cover_path
) VALUES
(
  '22222222-2222-2222-2222-222222222222',
  'Sample Album',
  2026,
  'covers/sample-album.jpg'
);

INSERT OR IGNORE INTO album_artists (
  album_id,
  artist_id,
  position
) VALUES
(
  '22222222-2222-2222-2222-222222222222',
  '11111111-1111-1111-1111-111111111111',
  0
);

INSERT OR IGNORE INTO tracks (
  id,
  album_id,
  title,
  disc_number,
  track_number,
  duration_ms,
  file_path,
  file_size,
  codec,
  bit_depth,
  sample_rate,
  bitrate,
  is_lossless,
  is_hires
) VALUES
(
  '33333333-3333-3333-3333-333333333333',
  '22222222-2222-2222-2222-222222222222',
  'Sample Hi-Res Track',
  1,
  1,
  210000,
  'dev-music/sample-hires.flac',
  52428800,
  'flac',
  24,
  192000,
  4608000,
  1,
  1
),
(
  '44444444-4444-4444-4444-444444444444',
  '22222222-2222-2222-2222-222222222222',
  'Sample Lossless Track',
  1,
  2,
  195000,
  'dev-music/sample-lossless.flac',
  31457280,
  'flac',
  16,
  44100,
  1411200,
  1,
  0
);

INSERT OR IGNORE INTO track_artists (
  track_id,
  artist_id,
  role,
  position
) VALUES
(
  '33333333-3333-3333-3333-333333333333',
  '11111111-1111-1111-1111-111111111111',
  'main',
  0
),
(
  '44444444-4444-4444-4444-444444444444',
  '11111111-1111-1111-1111-111111111111',
  'main',
  0
);

INSERT OR IGNORE INTO lyrics (
  id,
  track_id,
  format,
  language,
  content,
  offset_ms
) VALUES
(
  '55555555-5555-5555-5555-555555555555',
  '33333333-3333-3333-3333-333333333333',
  'lrc',
  'ja',
  '[00:00.00]Sample Hi-Res Track
[00:10.00]This is a demo lyric line
[00:20.00]TuneDock sample lyrics',
  0
);

INSERT OR IGNORE INTO search_index (
  track_id,
  title_search,
  artist_search,
  album_search,
  album_artist_search,
  filename_search
) VALUES
(
  '33333333-3333-3333-3333-333333333333',
  'sample hires track',
  'sample artist',
  'sample album',
  'sample artist',
  'sample-hires.flac'
),
(
  '44444444-4444-4444-4444-444444444444',
  'sample lossless track',
  'sample artist',
  'sample album',
  'sample artist',
  'sample-lossless.flac'
);
