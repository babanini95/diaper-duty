-- name: InsertDiaperChange :one
INSERT INTO changes (change_time, notes)
VALUES (?, ?)
RETURNING *;

-- name: GetTheLastChange :one
SELECT *
FROM changes
ORDER BY id DESC
LIMIT 1;