-- name: List :many
SELECT * FROM "hospital" WHERE "deletedAt" IS NULL;

-- name: Save :one
INSERT INTO "hospital" (
  "createdAt", "updatedAt", id, name, phone, city, "roadAddress", latitude, longitude, "zipCode"
) VALUES (
  NOW(), NOW(), $1, $2, $3, $4, $5, $6, $7, $8
) ON CONFLICT ("id") DO UPDATE
SET 
  "updatedAt" = NOW(),
  name = EXCLUDED.name, 
  phone = EXCLUDED.phone, 
  city = EXCLUDED.city,
  "roadAddress" = EXCLUDED."roadAddress",
  latitude = EXCLUDED.latitude,
  longitude = EXCLUDED.longitude,
  "zipCode" = EXCLUDED."zipCode"
RETURNING *;

-- name: Delete :exec
UPDATE "hospital" SET "deletedAt" = NOW()
WHERE id = $1;

-- name: FindByCity :many
SELECT * FROM hospital 
WHERE city = ANY($1::text[]);

-- name: CountByCity :one
SELECT COUNT(*) FROM hospital 
WHERE city = ANY($1::text[]);