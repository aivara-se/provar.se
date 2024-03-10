-- Create private.organization table
CREATE TABLE private.organization (
    id VARCHAR(8) PRIMARY KEY,
    slug VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(8) REFERENCES private.user(id),
    modified_at TIMESTAMP NOT NULL,
    name VARCHAR(64) NOT NULL
);

-- Index on slug column of private.organization table
CREATE INDEX idx_permission_slug ON private.organization(slug);

-- Create public.organization view
CREATE OR REPLACE VIEW public.organization AS
SELECT * FROM private.organization;
