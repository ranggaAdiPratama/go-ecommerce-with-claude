-- name: GetUserById :one
SELECT *
FROM users
WHERE
    id = $1
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE
    username = $1
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE
    email = $1
    AND deleted_at IS NULL
LIMIT 1;

-- name: UserList :many
SELECT * FROM users
WHERE
    deleted_at IS NULL
    AND (
        NULLIF(@role::text, '') IS NULL
        OR role = @role::text
    )
ORDER BY
    CASE
        WHEN @sort::text = 'name' AND @sort_order::text = 'asc' THEN name
    END ASC,
    CASE
        WHEN @sort::text = 'name' AND @sort_order::text = 'desc' THEN name
    END DESC,
    CASE
        WHEN @sort::text = 'created_at' AND @sort_order::text = 'asc' THEN created_at
    END ASC,
    CASE
        WHEN @sort::text = 'created_at' AND @sort_order::text = 'desc' THEN created_at
    END DESC
LIMIT COALESCE(@till::int, 15);

-- name: StoreUser :one
INSERT INTO
    users (
        name,
        username,
        email,
        password,
        role,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        NOw(),
        NOW()
    ) RETURNING id,
    name,
    username,
    email,
    role,
    created_at,
    updated_at;

-- name: UpdateUser :one
UPDATE users
SET
    name = COALESCE(NULLIF(@name::text, ''), name),
    username = COALESCE(
        NULLIF(@username::text, ''),
        username
    ),
    email = COALESCE(NULLIF(@email::text, ''), email),
    role = COALESCE(NULLIF(@role::text, ''), role),
    password = COALESCE(
        NULLIF(@password::text, ''),
        password
    ),
    updated_at = NOW()
WHERE
    id = @id RETURNING id,
    name,
    username,
    email,
    role,
    created_at,
    updated_at;

-- name: DeleteUser :exec
UPDATE users SET deleted_at = NOW() WHERE id = $1;