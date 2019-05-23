BEGIN;

--
-- Migration file for the `users` table. 
--

CREATE TABLE IF NOT EXISTS users (
    id              SERIAL,
    fullname        CHARACTER VARYING(255)          NOT NULL,
    city            CHARACTER VARYING(255)          NOT NULL,
    phone_number    CHARACTER VARYING(10) UNIQUE    NOT NULL,
    PRIMARY KEY(id)
);

COMMENT ON COLUMN users.id IS 'Autoincrement table ID';
COMMENT ON COLUMN users.fullname IS 'First name and last name of a user';
COMMENT ON COLUMN users.city IS 'City of a user';
COMMENT ON COLUMN users.phone_number IS 'Phone number of a user';


COMMIT;
