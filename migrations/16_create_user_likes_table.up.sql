CREATE TABLE user_likes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    music_id INTEGER REFERENCES music(id),
    post_id INTEGER REFERENCES posts(id),
    comment_id INTEGER REFERENCES comments(id),
    liked BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_likes_target_check CHECK (
        (music_id IS NOT NULL AND post_id IS NULL AND comment_id IS NULL) OR
        (music_id IS NULL AND post_id IS NOT NULL AND comment_id IS NULL) OR
        (music_id IS NULL AND post_id IS NULL AND comment_id IS NOT NULL)
    )
);