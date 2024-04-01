-- Create private.feedback table
CREATE TABLE private.feedback (
    id VARCHAR(8) PRIMARY KEY,
    organization_id VARCHAR(8) REFERENCES private.organization(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    question_type VARCHAR(32) NOT NULL,
    feedback_type VARCHAR(32) NOT NULL,
    feedback_time TIMESTAMP NOT NULL,
    feedback_data JSONB NOT NULL,
    feedback_tags JSONB NOT NULL,
    feedback_meta JSONB NOT NULL,
    feedback_user JSONB NOT NULL
);

-- Index on organization_id and created_at columns of private.feedback table
CREATE INDEX idx_feedback_organization_id_created_at ON private.feedback(organization_id, created_at);
