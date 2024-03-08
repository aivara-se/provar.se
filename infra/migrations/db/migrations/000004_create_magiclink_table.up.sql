-- Create "magiclink" table
CREATE TABLE magiclink (
    id VARCHAR(8) PRIMARY KEY,
    user_id VARCHAR(8) REFERENCES "user"(id),
    token VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index on token column of "magiclink" table
CREATE INDEX idx_magiclink_token ON magiclink(token);
