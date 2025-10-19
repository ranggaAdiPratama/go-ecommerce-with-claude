-- name: StoreRefreshToken :one
INSERT INTO
    refresh_tokens (
        user_id,
        token_hash,
        expires_at,
        created_at
    )
VALUES ($1, $2, $3, NOW()) RETURNING id,
    user_id,
    token_hash,
    is_revoked,
    expires_at,
    created_at,
    revoked_at;

-- name: GetRefreshToken :one
SELECT
    id,
    user_id,
    token_hash,
    is_revoked,
    expires_at,
    created_at,
    revoked_at
FROM refresh_tokens
WHERE
    token_hash = $1
LIMIT 1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET
    is_revoked = TRUE,
    revoked_at = NOW()
WHERE
    token_hash = $1;

-- name: RevokeAllUserRefreshTokens :exec
UPDATE refresh_tokens
SET
    is_revoked = TRUE,
    revoked_at = NOW()
WHERE
    user_id = $1
    AND is_revoked = FALSE;

-- name: DeleteExpiredRefreshTokens :exec
DELETE FROM refresh_tokens WHERE expires_at < NOW();

-- name: DeleteRevokedRefreshTokens :exec
DELETE FROM refresh_tokens
WHERE
    is_revoked = TRUE
    AND revoked_at < NOW() - INTERVAL '30 days';