-- name: GetDeletedShopById :one
SELECT *
FROM shops
WHERE
    id = $1
    AND deleted_at IS NOT NULL
LIMIT 1;

-- name: GetShopById :one
SELECT *
FROM shops
WHERE
    id = $1
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetShopByName :one
SELECT *
FROM shops
WHERE
    name = $1
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetShopByUserId :one
SELECT *
FROM shops
WHERE
    user_id = $1
    AND deleted_at IS NULL
LIMIT 1;

-- name: ShopList :many
SELECT * FROM shops
WHERE
    deleted_at IS NULL
     AND (
        NULLIF(@rank::text, '') IS NULL
        OR rank = @rank::text
    )
    AND (
        NULLIF(@search::text, '') IS NULL
        OR name ~* @search::text
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
    END DESC,
    name DESC
LIMIT COALESCE(@till::int, 15)
OFFSET COALESCE(@page::int, 0);

-- name: ShopListTotal :one
SELECT COUNT(*) AS total
FROM shops
WHERE
    deleted_at IS NULL
     AND (
        NULLIF(@rank::text, '') IS NULL
        OR rank = @rank::text
    )
    AND (
        NULLIF(@search::text, '') IS NULL
        OR name ~* @search::text
    );

-- name: StoreShop :one
INSERT INTO
    shops (
        user_id,
        name,
        logo,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, NOw(), NOW()) RETURNING id,
    user_id,
    name,
    logo,
    rank,
    created_at,
    updated_at;

-- name: UpdateShop :one
UPDATE shops
SET
    name = COALESCE(NULLIF(@name::text, ''), name),
    logo = COALESCE(
        NULLIF(@logo::text, ''),
        logo
    ),
    rank = COALESCE(
        NULLIF(@rank::text, ''),
        rank
    ),
    updated_at = NOW()
WHERE
    id = @id RETURNING id,
    user_id,
    name,
    logo,
    rank,
    created_at,
    updated_at;

-- name: DeleteShop :exec
UPDATE shops SET deleted_at = NOW() WHERE id = $1;