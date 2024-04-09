CREATE TABLE Categories
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE users
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE posts (
    id SERIAL NOT NULL UNIQUE,
    category_id INT  REFERENCES Categories(id) NOT NULL,
    user_id INT  REFERENCES users(id) NOT NULL,
    category_name VARCHAR(255),
    user_name VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
    image_path VARCHAR(255)
);



