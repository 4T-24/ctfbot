-- migrate:up
CREATE TABLE ctf (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    ctf_time_id INT NOT NULL
);

CREATE TABLE settings (
    k TEXT NOT NULL PRIMARY KEY,
    v TEXT NOT NULL
);

-- migrate:down
DROP TABLE settings;
DROP TABLE ctf;
