CREATE TABLE "shortened_urls" (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(16) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- owner INT,
    -- CONSTRAINT fk_owner FOREIGN KEY (owner)
    --     REFERENCES "user"(id)
    --     ON DELETE SET NULL
);
