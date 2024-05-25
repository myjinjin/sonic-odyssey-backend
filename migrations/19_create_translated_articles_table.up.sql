CREATE TABLE translated_articles (
    id SERIAL PRIMARY KEY,
    original_url VARCHAR(255) NOT NULL,
    translated_title VARCHAR(255) NOT NULL,
    translated_content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);