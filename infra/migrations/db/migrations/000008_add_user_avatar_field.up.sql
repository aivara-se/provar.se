-- Add the "avatar" column on private.user table
ALTER TABLE private.user
    ADD COLUMN avatar VARCHAR(128) DEFAULT '';
