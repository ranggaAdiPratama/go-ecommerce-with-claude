CREATE TABLE shops (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    logo VARCHAR(255) NOT NULL,
    rank VARCHAR(255) NOT NULL DEFAULT 'bronze',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_shops_user_id ON shops (user_id);

CREATE INDEX idx_shops_name ON shops (name);

CREATE INDEX idx_shops_rank ON shops (rank);

CREATE INDEX idx_shops_deleted_at ON shops (deleted_at);

ALTER TABLE shops
ADD CONSTRAINT check_shop_rank CHECK (
    rank IN (
        'bronze',
        'silver',
        'gold',
        'platinum'
    )
);