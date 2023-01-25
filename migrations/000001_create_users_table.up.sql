CREATE TABLE IF NOT EXISTS public.users (
    id bigserial PRIMARY KEY NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    passwd varchar(100) NOT NULL,
    phone_number varchar(20) NOT NULL,
    is_active smallint NOT NULL DEFAULT 0,
    created_at bigint NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL,
    updated_at bigint NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL
);