CREATE TABLE IF NOT EXISTS public.access (
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    is_active smallint NOT NULL DEFAULT 0,
    created_at int64 NOT NULL DEFAULT 0,
    updated_at int64 NOT NULL DEFAULT 0
);