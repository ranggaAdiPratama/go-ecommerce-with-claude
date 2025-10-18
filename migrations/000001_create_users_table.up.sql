CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_users_email ON users (email);

CREATE INDEX idx_users_username ON users (username);

CREATE INDEX idx_users_created_at ON users (created_at);

ALTER TABLE users
ADD CONSTRAINT check_user_role CHECK (
    role IN ('user', 'admin', 'seller')
);

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
        'Admin',
        'admin',
        'admin@mail.com',
        '$2a$10$kQ4DeLpy9YV7a1v4W3hu8OaFnRKXqv7uSlTeOKZEHGfImVStyguqC',
        'admin',
        NOw(),
        NOW()
    ),
    (
        'Rangga Adi Pratama',
        'rangga',
        'masterrangga@gmail.com',
        '$2a$10$N3aTkMc/FY9Ij76mreKqB.26udu20ryqaLdbfO2DFaOlmD.5bfEgy',
        'user',
        NOw(),
        NOW()
    ),
    (
        'Mitsuha Miyamizu',
        'mitsuha',
        'mitsuha@mail.com',
        '$2a$10$kQ4DeLpy9YV7a1v4W3hu8OaFnRKXqv7uSlTeOKZEHGfImVStyguqC',
        'seller',
        NOw(),
        NOW()
    );