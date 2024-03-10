-- Create "credential" table
CREATE TABLE credential (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    last_used_at TIMESTAMP,
    organization_id VARCHAR(8) REFERENCES organization(id),
    name VARCHAR(64) NOT NULL,
    secret VARCHAR(64) NOT NULL
);

-- Index on secret column of "credential" table
CREATE INDEX idx_credential_secret ON credential(secret);

-- Index on organization_id column of "credential" table
CREATE INDEX idx_credential_organization_id ON credential(organization_id);
