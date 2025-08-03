-- +goose Up
CREATE TABLE profiles (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    baby_name TEXT NOT NULL,
    baby_birthday TEXT NOT NULL,
    diaper_interval_minutes INTEGER
);

CREATE TABLE changes (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    change_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

-- +goose Down
DROP TABLE profiles;

DROP TABLE changes;