CREATE TABLE IF NOT EXISTS public.articles (
    id bigserial PRIMARY KEY NOT NULL,
    category_id bigserial NOT NULL,
    title varchar(100) NOT NULL,
    content text,
    status smallint NOT NULL DEFAULT 0,
    created_at int64 NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL,
    updated_at int64 NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (id),
    CONSTRAINT fk_user FOREIGN KEY (created_by, updated_by) REFERENCES users (id)  
);