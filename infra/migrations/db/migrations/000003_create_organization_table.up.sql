-- Create private.organization table
CREATE TABLE private.organization (
    id VARCHAR(8) PRIMARY KEY,
    slug VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(8) REFERENCES private.user(id),
    modified_at TIMESTAMP NOT NULL,
    name VARCHAR(64) NOT NULL,
    description VARCHAR(128) DEFAULT ''
);

-- Index on slug column of private.organization table
CREATE INDEX idx_permission_slug ON private.organization(slug);
