-- Create private.invitation table
CREATE TABLE private.invitation (
    id VARCHAR(8) PRIMARY KEY,
    organization_id VARCHAR(8) REFERENCES private.organization(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(8) REFERENCES private.user(id) ON DELETE CASCADE,
    expires_at TIMESTAMP NOT NULL,
    accepted_at TIMESTAMP,
    secret VARCHAR(64) NOT NULL,
    email VARCHAR(64) UNIQUE NOT NULL,
    name VARCHAR(64) NOT NULL
);

-- Index on secret column of private.invitation table
CREATE INDEX idx_invitation_secret ON private.invitation(secret);

-- Index on organization_id column of private.invitation table
CREATE INDEX idx_invitation_organization_id ON private.invitation(organization_id);
