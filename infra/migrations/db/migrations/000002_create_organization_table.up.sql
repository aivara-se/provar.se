-- Create "organization" table
CREATE TABLE organization (
    id VARCHAR(8) PRIMARY KEY,
    slug VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    owner_id VARCHAR(8) REFERENCES "user"(id),
    name VARCHAR(64) NOT NULL
);
