CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email varchar(255) UNIQUE NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255),
    password varchar(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT
);
