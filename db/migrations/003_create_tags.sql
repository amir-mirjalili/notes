-- +migrate Up

CREATE TABLE IF NOT EXISTS tags (
                      id SERIAL PRIMARY KEY,
                      name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS note_tags (
                           note_id INT REFERENCES notes(id),
                           tag_id INT REFERENCES tags(id),
                           PRIMARY KEY (note_id, tag_id)
);

-- +migrate Down
DROP TABLE tags;
DROP TABLE note_tags;
