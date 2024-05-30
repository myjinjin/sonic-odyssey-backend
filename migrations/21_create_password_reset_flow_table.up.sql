CREATE TABLE password_reset_flows (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    flow_id VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_password_reset_flows_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);