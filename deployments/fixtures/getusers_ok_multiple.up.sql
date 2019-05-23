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
    ('Sean Case', 'Paris', '0625000010'),
    ('Roberto Polo', 'Montpellier', '0625000011'),
    ('Lucas Robert', 'Vitry-sur-Seine', '0625000012');

COMMIT;
