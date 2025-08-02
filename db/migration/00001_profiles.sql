-- +goose Up
CREATE TABLE profiles (
    id INTEGER PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    baby_name TEXT NOT NULL,
    baby_birthday TEXT NOT NULL,
    diaper_interval_minutes INTEGER
);

CREATE TABLE changes (
    id INTEGER PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    change_time TEXT NOT NULL,
    notes TEXT
);

-- +goose Down
DROP TABLE profiles;

DROP TABLE changes;