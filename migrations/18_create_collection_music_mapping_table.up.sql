CREATE TABLE collection_music_mapping (
    id SERIAL PRIMARY KEY,
    collection_id INTEGER REFERENCES music_collections(id),
    music_id INTEGER REFERENCES music(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);