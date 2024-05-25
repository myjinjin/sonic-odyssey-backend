CREATE TABLE music_artist_mapping (
    id SERIAL PRIMARY KEY,
    music_id INTEGER REFERENCES music(id),
    artist_id INTEGER REFERENCES artists(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);