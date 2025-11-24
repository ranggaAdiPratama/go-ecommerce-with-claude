CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) UNIQUE NOT NULL,
    icon VARCHAR(100) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_categories_slug ON categories (slug);

CREATE INDEX idx_categories_is_active ON categories (is_active);

INSERT INTO
    categories (
        name,
        icon,
        slug,
        is_active,
        created_at,
        updated_at
    )
VALUES (
        'Electronics',
        'Laptop',
        'electronics',
        true,
        NOw(),
        NOW()
    ),
    (
        'Fashion',
        'Shirt',
        'fashion',
        true,
        NOw(),
        NOW()
    );