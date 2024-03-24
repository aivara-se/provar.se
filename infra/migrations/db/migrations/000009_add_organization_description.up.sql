-- Add the "description" column on private.organization table
ALTER TABLE private.organization
    ADD COLUMN description VARCHAR(128) DEFAULT '';
