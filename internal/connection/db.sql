CREATE TABLE gomigrate_migrators (
    id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Version INT NOT NULL DEFAULT 1,
    CreatedAt BIGINT NOT NULL
);

CREATE TABLE gomigrate_migrator_stash (
    id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Version INT NOT NULL,
    upFileBody TEXT,
    downFileBody TEXT,
    ChangedAt BIGINT NOT NULL
);

INSERT INTO gomigrate_migrators (Name, CreatedAt) VALUES ('%s', '%s') RETURNING *;