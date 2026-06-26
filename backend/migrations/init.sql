PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS artists (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  sort_name TEXT,
  cover_id  TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS albums (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  year INTEGER,
  cover_id TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS album_artists (
  album_id TEXT NOT NULL,
  artist_id TEXT NOT NULL,
  position INTEGER DEFAULT 0,

  PRIMARY KEY (album_id, artist_id),
  FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE,
  FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tracks (
  id TEXT PRIMARY KEY,
  album_id TEXT,
  title TEXT NOT NULL,

  disc_number INTEGER DEFAULT 1,
  track_number INTEGER,
  duration_ms INTEGER,

  file_path TEXT NOT NULL UNIQUE,
  file_size INTEGER,

  cover_id  TEXT,

  codec TEXT,
  bit_depth INTEGER,
  sample_rate INTEGER,
  bitrate INTEGER,

  is_lossless INTEGER DEFAULT 0,
  is_hires INTEGER DEFAULT 0,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS track_artists (
  track_id TEXT NOT NULL,
  artist_id TEXT NOT NULL,
  role TEXT DEFAULT 'main',
  position INTEGER DEFAULT 0,

  PRIMARY KEY (track_id, artist_id, role),
  FOREIGN KEY (track_id) REFERENCES tracks(id) ON DELETE CASCADE,
  FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lyrics (
  id TEXT PRIMARY KEY,
  track_id TEXT NOT NULL,
  format TEXT NOT NULL,
  language TEXT,
  content TEXT NOT NULL,
  offset_ms INTEGER DEFAULT 0,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (track_id) REFERENCES tracks(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS search_index (
  track_id TEXT PRIMARY KEY,
  title_search TEXT,
  artist_search TEXT,
  album_search TEXT,
  album_artist_search TEXT,
  filename_search TEXT,

  FOREIGN KEY (track_id) REFERENCES tracks(id) ON DELETE CASCADE
);
