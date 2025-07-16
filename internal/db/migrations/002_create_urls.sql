CREATE TABLE "shortened_urls" (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(16) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    owner_id BIGINT NOT NULL,
    CONSTRAINT fk_shortened_urls_owner FOREIGN KEY (owner_id)
        REFERENCES "users"(id)
        ON DELETE CASCADE
);

