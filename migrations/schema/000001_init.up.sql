BEGIN;
    CREATE TABLE users (
        id serial not null unique,
        login varchar (255) not null unique,
        password_hash varchar(255) not null
    );
COMMIT;