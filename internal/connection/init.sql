CREATE TABLE gomigrate_migrators (
    id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Version INT NOT NULL DEFAULT 1,
    CreatedAt BIGINT NOT NULL,
    ChangedAt BIGINT NOT NULL
);

CREATE TABLE gomigrate_migrators_stash (
    id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Version INT NOT NULL,
    upFileBody TEXT,
    downFileBody TEXT,
    ChangedAt BIGINT NOT NULL
);