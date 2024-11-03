-- name: FindByValue :one
SELECT
    *
FROM
    "refreshToken"
WHERE
    "value" = $1;

-- name: List :many
SELECT
    *
FROM
    "refreshToken" LIMIT $1 OFFSET $2;

-- name: Save :one
INSERT INTO
    "refreshToken" (
        "createdAt",
        "updatedAt",
        "userId",
        "value",
        "clientInfo"
    )
VALUES
    (NOW(), NOW(), $1, $2, $3) RETURNING *;

-- name: Delete :exec
DELETE FROM
    "refreshToken"
WHERE
    id = $1;