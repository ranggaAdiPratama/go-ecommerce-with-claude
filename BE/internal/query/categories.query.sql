-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: GetCategoryByName :one
SELECT * FROM categories WHERE name = $1 LIMIT 1;

-- name: GetCategoryBySlug :one
SELECT *
FROM categories
WHERE
    slug = $1
    AND is_active = true
LIMIT 1;

-- name: CategoryList :many
SELECT * FROM categories
WHERE 
    (
        NULLIF(@search::text, '') IS NULL
        OR name ~* @search::text
    )
    AND (
        (NULLIF(@status::text, '') IS NULL AND is_active = TRUE)
        OR (@status::text = '1' AND is_active = TRUE)
        OR (@status::text = '0' AND is_active = FALSE)
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

-- name: CategoryListTotal :one
SELECT COUNT(*) AS total
FROM categories
WHERE
    (
        NULLIF(@search::text, '') IS NULL
        OR name ~* @search::text
    )
    AND (
        NULLIF(@status::text, '') IS NULL
        OR (
            (@status::text = '1' AND is_active = TRUE)
            OR (@status::text = '0' AND is_active = FALSE)
        )
    );

-- name: StoreCategory :one
INSERT INTO
    categories (
        name,
        icon,
        slug,
        is_active,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, NOw(), NOW()) RETURNING id,
    name,
    icon,
    slug,
    is_active,
    created_at,
    updated_at;

-- name: UpdateCategory :one
UPDATE categories
SET
    name = COALESCE(NULLIF(@name::text, ''), name),
    icon = COALESCE(
        NULLIF(@icon::text, ''),
        icon
    ),
    slug = @slug::text,
    updated_at = NOW()
WHERE
    id = @id RETURNING id,
    name,
    icon,
    slug,
    is_active,
    created_at,
    updated_at;

-- name: SwitchOn :one
UPDATE categories
SET
    is_active = TRUE,
    updated_at = NOW()
WHERE
    id = @id RETURNING id,
    name,
    icon,
    slug,
    is_active,
    created_at,
    updated_at;

-- name: SwitchOff :one
UPDATE categories
SET
    is_active = FALSE,
    updated_at = NOW()
WHERE
    id = @id RETURNING id,
    name,
    icon,
    slug,
    is_active,
    created_at,
    updated_at;