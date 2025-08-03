-- name: CountProfiles :one
SELECT count(*)
FROM profiles;

-- name: CreateProfile :one
INSERT INTO profiles (baby_name, baby_birthday)
VALUES (?, ?)
RETURNING *;