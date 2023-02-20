CREATE DATABASE bapi;

CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   email VARCHAR(255) NOT NULL,
   username VARCHAR(255) NOT NULL,
   is_activated Boolean DEFAULT FALSE,
   password TEXT NOT NULL
);

CREATE TABLE tokens (
    token_id INT REFERENCES users(id),
    UNIQUE(token_id),
    refreshToken TEXT NOT NULL
);