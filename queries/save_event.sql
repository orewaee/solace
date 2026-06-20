-- name: InsertEvent :execresult
INSERT INTO events (id, user_id, text, start_at, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);
