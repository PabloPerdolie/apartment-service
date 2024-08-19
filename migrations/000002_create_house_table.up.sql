CREATE TABLE IF NOT EXISTS houses (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL UNIQUE,
    year INT NOT NULL,
    developer VARCHAR(50),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
