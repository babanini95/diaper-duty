-- name: CreateProfilesTable :exec
CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    baby_name TEXT NOT NULL,
    baby_birthday TEXT NOT NULL,
    diaper_interval_minutes INTEGER
);

-- name: CreateChangesTable :exec
CREATE TABLE IF NOT EXISTS changes (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    change_time TEXT NOT NULL,
    notes TEXT
);