CREATE TABLE todos
(
    id          SERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    user_id     INTEGER      NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);