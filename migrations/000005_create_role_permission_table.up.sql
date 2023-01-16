CREATE TABLE IF NOT EXISTS public.role_permissions (
    role_id bigserial,
    permission_id bigserial,
    created_at int64 DEFAULT 0, 
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
);