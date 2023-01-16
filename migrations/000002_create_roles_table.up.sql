CREATE TABLE IF NOT EXISTS public.roles (
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    is_active smallint NOT NULL,
    created_at int64 NOT NULL DEFAULT 0,
    updated_at int64 NOT NULL DEFAULT 0
);