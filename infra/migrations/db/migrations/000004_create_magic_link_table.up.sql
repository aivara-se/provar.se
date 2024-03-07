-- Create "magic_link" table
CREATE TABLE magic_link (
    id VARCHAR(8) PRIMARY KEY,
    user_id VARCHAR(8) REFERENCES "user"(id),
    token VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index on token column of "magic_link" table
CREATE INDEX idx_magic_link_token ON magic_link(token);
