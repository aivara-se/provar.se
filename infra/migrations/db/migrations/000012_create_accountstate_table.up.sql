-- Create private.accountstate table
CREATE TABLE private.accountstate (
    id VARCHAR(8) PRIMARY KEY,
    account_id VARCHAR(8) REFERENCES private.account(id) ON DELETE CASCADE,
    state VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index on account_id column of private.accountstate table
CREATE INDEX idx_accountstate_account_id ON private.accountstate(account_id);

-- Index on state column of private.accountstate table
CREATE INDEX idx_accountstate_state ON private.accountstate(state);
