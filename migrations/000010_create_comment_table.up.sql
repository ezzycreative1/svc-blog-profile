CREATE TABLE IF NOT EXISTS public.comments (
    id bigserial PRIMARY KEY NOT NULL,
    article_id bigserial NOT NULL,
    user_id bigserial NOT NULL,
    title varchar(100) NOT NULL,
    comment text,
    status smallint NOT NULL DEFAULT 0,
    created_at bigint NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    updated_at bigint NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE
);