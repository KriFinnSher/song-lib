CREATE TABLE IF NOT EXISTS songs (
                       id SERIAL PRIMARY KEY,
                       artist VARCHAR(255) NOT NULL,
                       title VARCHAR(255) NOT NULL,
                       release_date VARCHAR(255) NOT NULL,
                       text TEXT NOT NULL,
                       source_link TEXT NOT NULL,
                       UNIQUE (artist, title)
);
