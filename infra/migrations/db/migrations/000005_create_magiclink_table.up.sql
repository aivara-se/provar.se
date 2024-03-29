-- Create private.magiclink table
CREATE TABLE private.magiclink (
    id VARCHAR(8) PRIMARY KEY,
    user_id VARCHAR(8) REFERENCES private.user(id) ON DELETE CASCADE,
    token VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index on token column of private.magiclink table
CREATE INDEX idx_magiclink_token ON private.magiclink(token);
