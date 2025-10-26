-- name: GetSetting :one
SELECT * FROM settings LIMIT 1;

-- name: StoreSetting :one
INSERT INTO
    settings (
        name,
        logo,
        created_at,
        updated_at
    )
VALUES ($1, $2, NOW(), NOW()) RETURNING name,
    logo;

-- name: UpdateSetting :one
UPDATE settings
SET
    name = COALESCE(NULLIF(@name::text, ''), name),
    logo = COALESCE(
        NULLIF(@logo::text, ''),
        logo
    ),
    updated_at = NOW()
WHERE
    id = @id RETURNING name,
    logo;