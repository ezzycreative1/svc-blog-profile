CREATE TABLE IF NOT EXISTS public.user_roles (
    user_id bigserial,
    role_id bigserial,
    created_at bigint DEFAULT 0,
    verify_token varchar(200),
    verify_at bigint DEFAULT 0,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);