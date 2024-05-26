DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'social_provider') THEN
        CREATE TYPE social_provider AS ENUM ('GOOGLE', 'KAKAO', 'NAVER', 'SPOTIFY');
    END IF;
END$$;

CREATE TABLE user_social_accounts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    provider social_provider NOT NULL,
    provider_user_id VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);