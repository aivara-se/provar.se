-- Create private.credential table
CREATE TABLE private.credential (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(8) REFERENCES private.user(id) ON DELETE CASCADE,
    modified_at TIMESTAMP NOT NULL,
    last_used_at TIMESTAMP,
    organization_id VARCHAR(8) REFERENCES private.organization(id) ON DELETE CASCADE,
    name VARCHAR(64) NOT NULL,
    secret VARCHAR(64) NOT NULL
);

-- Index on secret column of private.credential table
CREATE INDEX idx_credential_secret ON private.credential(secret);

-- Index on organization_id column of private.credential table
CREATE INDEX idx_credential_organization_id ON private.credential(organization_id);
