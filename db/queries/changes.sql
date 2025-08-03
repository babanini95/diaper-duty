-- name: InsertDiaperChange :one
INSERT INTO changes (change_time, notes)
VALUES (?, ?)
RETURNING *;