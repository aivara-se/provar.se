-- Create "user" table
CREATE TABLE "user" (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    email VARCHAR(64) UNIQUE NOT NULL,
    email_verified_at TIMESTAMP,
    name VARCHAR(64) NOT NULL
);

-- Index on email column of "user" table
CREATE INDEX idx_user_email ON "user"(email);
