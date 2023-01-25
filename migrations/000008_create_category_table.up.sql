CREATE TABLE IF NOT EXISTS public.categories (
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    desctiption text,
    is_active smallint NOT NULL DEFAULT 0,
    created_at bigint NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL,
    updated_at bigint NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL
);