CREATE TABLE IF NOT EXISTS public.permissions (
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    server_name varchar(255) NOT NULL,
    path_name varchar(255) NOT NULL,
    is_active smallint NOT NULL,
    created_at bigint NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL,
    updated_at bigint NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL
);