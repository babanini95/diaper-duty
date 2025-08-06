-- name: CountProfiles :one
SELECT count(*)
FROM profiles;

-- name: CreateProfile :one
INSERT INTO profiles (baby_name, baby_birthday)
VALUES (?, ?)
RETURNING *;

-- name: GetProfile :one
SELECT *
FROM profiles
LIMIT 1;

-- name: SetCustomReminder :exec
UPDATE profiles
SET diaper_interval_minutes = ?
RETURNING *;

-- name: ResetCustomReminder :exec
UPDATE profiles
SET diaper_interval_minutes = NULL
RETURNING *;