
-- CREATE TABLE movies (
--     id SERIAL PRIMARY KEY,
--     title VARCHAR(255) NOT NULL,
--     overview TEXT,
--     release_date DATE,
--     runtime INT,
--     budget NUMERIC(15,2),
--     original_language VARCHAR(50),
--     trailer_yt VARCHAR(255)
-- );

CREATE TABLE IF NOT EXISTS movies (
    id bigint,
    title varchar(355),
    movie JSON,
    primary key (id)
);

