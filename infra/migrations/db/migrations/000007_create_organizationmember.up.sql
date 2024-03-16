-- Create private.organizationmember table
CREATE TABLE private.organizationmember (
    id VARCHAR(8) PRIMARY KEY,
    user_id VARCHAR(8) REFERENCES private.user(id),
    organization_id VARCHAR(8) REFERENCES private.organization(id)
);

-- Index on user_id column of private.organizationmember table
CREATE INDEX idx_organizationmember_user_id ON private.organizationmember(user_id);

-- Index on organization_id column of private.organizationmember table
CREATE INDEX idx_organizationmember_organization_id ON private.organizationmember(organization_id);

-- Create public.organizationmember view
CREATE OR REPLACE VIEW public.organizationmember AS
SELECT * FROM private.organizationmember;
