CREATE TABLE users (
    id   SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password  TEXT NOT NULL
    -- created_at DATE,
    -- updated_at DATE
);

CREATE TABLE lists (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    -- created_at DATE,
    -- updated_at DATE,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE items (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    date DATE,
    -- created_at DATE,
    -- updated_at DATE,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

