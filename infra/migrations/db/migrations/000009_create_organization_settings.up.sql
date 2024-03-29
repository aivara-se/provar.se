-- Create private.organizationsetting table
CREATE TABLE private.organizationsetting (
    id VARCHAR(8) PRIMARY KEY,
    organization_id VARCHAR(8) REFERENCES private.organization(id) ON DELETE CASCADE,
    key VARCHAR(128) NOT NULL,
    val VARCHAR(128) NOT NULL
);

-- Index on organization_id column of private.organizationsetting table
CREATE INDEX idx_organizationsetting_organization_id ON private.organizationsetting(organization_id);
