INSERT INTO users 
VALUES 
    (1, 'abc@gmail.com', now(), now(), NULL),
    (2, 'def@gmmail.com', now(), now(), NULL);

INSERT INTO relationships (user_id_1, user_id_2, relationship_type, created_at, updated_at)
VALUES
    (1, 2, 'friend', now(), now());