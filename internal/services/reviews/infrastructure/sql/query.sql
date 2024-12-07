-- name: Save :one
INSERT INTO "review" (
  "createdAt", "updatedAt", id, "userId", "hospitalId", "content", "rating"
) VALUES (
  NOW(), NOW(), $1, $2, $3, $4, $5
) ON CONFLICT ("id") DO UPDATE
SET 
  "updatedAt" = NOW(),
  "content" = EXCLUDED."content", 
  "rating"=EXCLUDED."rating"
RETURNING *;

-- name: FindOne :one
SELECT * FROM "review" WHERE id = $1;

-- name: FindByHospitalId :many
SELECT * FROM "review" WHERE "hospitalId" = $1;

-- name: FindByUserIdAndHospitalId :one
SELECT * FROM "review" WHERE "userId" = $1 AND "hospitalId" = $2 LIMIT 1;

