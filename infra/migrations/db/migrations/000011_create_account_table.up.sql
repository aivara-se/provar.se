-- Create private.account table
CREATE TABLE private.account (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id VARCHAR(8) REFERENCES private.user(id) ON DELETE CASCADE,
    provider VARCHAR(16) NOT NULL,
    session VARCHAR(2048) NOT NULL,
    name VARCHAR(64) NOT NULL,
    avatar VARCHAR(128) DEFAULT ''
);
