-- Remove the "description" column from private.organization table
ALTER TABLE private.organization
    DROP COLUMN description;
