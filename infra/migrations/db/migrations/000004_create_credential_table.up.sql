-- Create private.credential table
CREATE TABLE private.credential (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    last_used_at TIMESTAMP,
    organization_id VARCHAR(8) REFERENCES private.organization(id),
    name VARCHAR(64) NOT NULL,
    secret VARCHAR(64) NOT NULL
);

-- Index on secret column of private.credential table
CREATE INDEX idx_credential_secret ON private.credential(secret);

-- Index on organization_id column of private.credential table
CREATE INDEX idx_credential_organization_id ON private.credential(organization_id);

-- Create public.credential view
CREATE OR REPLACE VIEW public.credential AS
SELECT * FROM private.credential;
