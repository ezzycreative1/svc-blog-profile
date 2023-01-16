CREATE TABLE IF NOT EXISTS public.permission_access (
    permission_id bigserial,
    access_id bigserial,
    created_at int64 DEFAULT 0, 
    FOREIGN KEY (permission_id) REFERENCES permissions (id),
    FOREIGN KEY (access_id) REFERENCES access (id)
);