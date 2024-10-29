-- name: FindOne :one
SELECT * FROM "user" WHERE id = $1;

-- name: List :many
SELECT * FROM "user" WHERE "deletedAt" IS NULL;

-- name: Save :one
INSERT INTO "user" (
  "createdAt", "updatedAt", id, email, nickname, "profileImageUrl", providers,"lastProviderType"
) VALUES (
  NOW(), NOW(), $1, $2, $3, $4 ,$5, $6
) ON CONFLICT ("email") DO UPDATE
SET 
  "updatedAt" = NOW(),
  nickname = EXCLUDED.nickname, 
  "profileImageUrl"=EXCLUDED."profileImageUrl", 
  providers = EXCLUDED.providers,
  "lastProviderType"=EXCLUDED."lastProviderType"
RETURNING *;

-- name: Delete :exec
UPDATE "user" SET "deletedAt" = NOW()
WHERE id = $1;

-- name: FindByEmail :one
SELECT * FROM "user" WHERE email = $1;