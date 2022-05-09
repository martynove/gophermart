BEGIN;
    CREATE TABLE users (
        id serial not null unique,
        login varchar (255) not null unique,
        password_hash varchar(255) not null
    );
    CREATE TABLE orders (
        number varchar (255) not null unique,
        user_id int references users (id) not null,
        status varchar( 255),
        accrual decimal,
        uploaded_at timestamp with time zone not null
    );
COMMIT;