-- Remove the "avatar" column from private.user table
ALTER TABLE private.user
    DROP COLUMN avatar;
