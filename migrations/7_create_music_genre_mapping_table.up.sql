CREATE TABLE music_genre_mapping (
    id SERIAL PRIMARY KEY,
    music_id INTEGER REFERENCES music(id),
    genre_id INTEGER REFERENCES genres(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);