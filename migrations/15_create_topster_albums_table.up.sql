CREATE TABLE topster_albums (
    id SERIAL PRIMARY KEY,
    topster_id INTEGER REFERENCES user_topsters(id),
    album_id INTEGER REFERENCES albums(id),
    position JSONB NOT NULL
);