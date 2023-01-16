CREATE TABLE IF NOT EXISTS public.tags (
    id bigserial PRIMARY KEY NOT NULL,
    article_id bigserial NOT NULL,
    user_id bigserial NOT NULL,
    name varchar(100) NOT NULL,
    status smallint NOT NULL DEFAULT 0,
    created_at int64 NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL,
    updated_at int64 NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL,
    CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES articles (id),
    CONSTRAINT fk_user FOREIGN KEY (created_by, updated_by) REFERENCES users (id)
);