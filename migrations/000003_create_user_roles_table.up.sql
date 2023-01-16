CREATE TABLE IF NOT EXISTS public.user_roles (
    user_id bigserial,
    role_id bigserial,
    verify_token string,
    verify_at int64 DEFAULT 0,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);