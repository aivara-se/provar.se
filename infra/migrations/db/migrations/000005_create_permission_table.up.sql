-- Create "permission" table
CREATE TABLE permission (
    id VARCHAR(8) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    organization_id VARCHAR(8) REFERENCES organization(id),
    principal_type VARCHAR(16) NOT NULL,
    principal VARCHAR(8) NOT NULL,
    resource_type VARCHAR(16) NOT NULL,
    resources VARCHAR(8) NOT NULL,
    permission VARCHAR(64) NOT NULL
);

-- Index on organization_id column of "permission" table
CREATE INDEX idx_permission_organization_id ON permission(organization_id);
