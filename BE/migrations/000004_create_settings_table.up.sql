CREATE TABLE settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    logo VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO
    settings (
        name,
        logo,
        created_at,
        updated_at
    )
VALUES (
        'Warunk Aink',
        'https://res.cloudinary.com/duwqriyoo/image/upload/v1761464684/go-e-commerce/1761464683_warunk%20aink.jpeg.jpg',
        NOw(),
        NOW()
    );
;