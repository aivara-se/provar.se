-- Create private.permission table
CREATE TABLE private.permission (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    organization_id VARCHAR(8) REFERENCES private.organization(id),
    principal_type VARCHAR(16) NOT NULL,
    principal VARCHAR(8) NOT NULL,
    resource_type VARCHAR(16) NOT NULL,
    resources VARCHAR(8) NOT NULL,
    permission VARCHAR(64) NOT NULL
);

-- Index on organization_id column of private.permission table
CREATE INDEX idx_permission_organization_id ON private.permission(organization_id);

-- Create public.permission view
CREATE OR REPLACE VIEW public.permission AS
SELECT * FROM private.permission;
