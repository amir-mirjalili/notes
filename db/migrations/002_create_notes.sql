-- +migrate Up

CREATE TABLE IF NOT EXISTS notes (
                       id SERIAL PRIMARY KEY,
                       user_id INT REFERENCES users(id),
                       title TEXT NOT NULL,
                       content TEXT,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE notes;