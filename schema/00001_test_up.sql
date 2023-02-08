CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) NOT NULL,
                       username VARCHAR(255) NOT NULL,
                       is_activated Boolean DEFAULT FALSE,
                       password TEXT NOT NULL
);

CREATE TABLE tokens (
                        token_id SERIAL PRIMARY KEY,
                        refreshToken TEXT NOT NULL
);

CREATE TABLE user_api (
                          id SERIAL PRIMARY KEY,
                          owner_id INT NOT NULL,
                          UNIQUE(owner_id),
                          api_name VARCHAR(255) NOT NULL,
                          api_key VARCHAR(255) NOT NULL,
                          createdAt timestamp DEFAULT CURRENT_TIMESTAMP,
                          expiryTime timestamp,
                          bearer VARCHAR(255) NOT NULL
);
