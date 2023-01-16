CREATE TABLE IF NOT EXISTS public.users (
    id bigserial PRIMARY KEY NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    passwd varchar(100) NOT NULL,
    phone_number varchar(20) NOT NULL,
    role_id bigserial NOT NULL DEFAULT 0,
    is_active smallint NOT NULL DEFAULT 0,
    created_at int64 NOT NULL DEFAULT 0,
    updated_at int64 NOT NULL DEFAULT 0,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles (id)
);