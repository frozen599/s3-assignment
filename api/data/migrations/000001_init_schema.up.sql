CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email varchar(256) NOT NULL UNIQUE,
    password_hash varchar(256) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


CREATE TABLE IF NOT EXISTS friends (
    id serial PRIMARY KEY,
    user_id_1 int NOT NULL,
    user_id_2 int NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
    CONSTRAINT unique_friends UNIQUE(user_id_1, user_id_2),
);

CREATE TABLE IF NOT EXISTS subscribers (
    id serial PRIMARY KEY,
    requestor_user_id int,
    target_user_id int,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

