BEGIN;

--
-- Fixture file for the `AddUser/ok_overwriting` test. 
--


INSERT INTO users (
    fullname,
    city,
    phone_number
)
VALUES
    ('Already Exist', 'Paris', '0625000021');

COMMIT;
