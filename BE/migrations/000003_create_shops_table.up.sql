CREATE TABLE shops (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    logo VARCHAR(255) NOT NULL,
    rank VARCHAR(255) NOT NULL DEFAULT 'bronze',
    slug VARCHAR(255) NOT NULL,
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

INSERT INTO
    shops (
        user_id,
        name,
        logo,
        slug,
        created_at,
        updated_at
    )
VALUES (
        (
            SELECT id
            FROM users
            WHERE
                username = 'mitsuha'
        ),
        'Ranmitsu Shop',
        'https://res.cloudinary.com/duwqriyoo/image/upload/v1761057333/go-e-commerce/1761057331_dreamina-2025-10-21-9882-a%20simple%20logo%20for%20a%20shop%20named%20Ranmitsu.....jpeg.jpg',
        'ranmitsu-shop',
        NOw(),
        NOW()
    ),
    (
        (
            SELECT id
            FROM users
            WHERE
                username = 'rangga'
        ),
        'Sumber Fortune',
        'https://res.cloudinary.com/duwqriyoo/image/upload/v1761277637/go-e-commerce/1761277638_sumber%20fortune.jpeg.jpg',
        'sumber-fortune',
        NOw(),
        NOW()
    );
;