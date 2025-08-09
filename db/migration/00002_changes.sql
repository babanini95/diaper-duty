-- +goose Up
CREATE TABLE IF NOT EXISTS changes (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    change_time TEXT NOT NULL,
    notes TEXT
);

-- +goose Down
DROP TABLE changes;