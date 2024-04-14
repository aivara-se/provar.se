-- Create private.kvstorage table
CREATE TABLE private.kvstorage (
    id VARCHAR(32) PRIMARY KEY,
    payload VARCHAR(2048) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);
