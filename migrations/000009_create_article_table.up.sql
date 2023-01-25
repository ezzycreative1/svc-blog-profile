CREATE TABLE IF NOT EXISTS public.articles (
    id bigserial PRIMARY KEY NOT NULL,
    category_id bigserial NOT NULL,
    title varchar(100) NOT NULL,
    content text,
    status smallint NOT NULL DEFAULT 0,
    created_at bigint NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    updated_at bigint NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);