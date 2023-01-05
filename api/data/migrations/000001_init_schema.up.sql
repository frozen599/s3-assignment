CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email text NOT NULL UNIQUE CHECK (email <> ''),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


CREATE TABLE IF NOT EXISTS relationships (
    id serial PRIMARY KEY,
    user_id_1 int NOT NULL,
    user_id_2 int NOT NULL,
    relationship_type TEXT NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT unique_relationship UNIQUE(user_id_1, user_id_2, relationship_type)
);