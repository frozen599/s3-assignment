INSERT INTO users 
VALUES 
    (1, 'abc@gmail.com', now(), now(), NULL),
    (2, 'def@gmail.com', now(), now(), NULL),
    (3, 'ghi@gmail.com', now(), now(), NULL);

INSERT INTO relationships
VALUES  
    (2, 1, 'subscriber', now(), now(), NULL);