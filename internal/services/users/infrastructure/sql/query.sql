-- name: FindOne :one
SELECT * FROM "user" WHERE id = $1;

-- name: List :many
SELECT * FROM "user" WHERE "deletedAt" IS NULL;

-- name: Create :one
INSERT INTO "user" (
  "createdAt", "updatedAt", id, email, nickname, providers
) VALUES (
  NOW(), NOW(), $1, $2, $3, $4
)
RETURNING *;

-- name: Delete :exec
UPDATE "user" SET "deletedAt" = NOW()
WHERE id = $1;