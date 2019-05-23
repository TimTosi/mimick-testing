BEGIN;

--
-- Fixture file for the `GetUsers/ok_one` test. 
--


INSERT INTO users (
    fullname,
    city,
    phone_number
)
VALUES
    ('Jeanne Dupont', 'Paris', '0625000000');

COMMIT;
