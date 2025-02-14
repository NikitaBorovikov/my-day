CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    enc_password TEXT NOT NULL,
    reg_date DATE 
);

CREATE TABLE task (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    is_important BOOLEAN,
    is_done BOOLEAN,
    due_date DATE,
    created_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE events (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    appointed_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
