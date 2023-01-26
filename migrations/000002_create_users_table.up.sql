CREATE TABLE IF NOT EXISTS public.users (
    id bigserial PRIMARY KEY NOT NULL,
    user_no varchar(255) NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    passwd varchar(100) NOT NULL,
    phone_number varchar(20) NOT NULL,
    role_id bigserial NOT NULL,
    verify_token varchar(200) NOT NULL,
    verify_at bigint NOT NULL DEFAULT 0,
    is_active smallint NOT NULL DEFAULT 0,
    created_at bigint NOT NULL DEFAULT 0,
    created_by bigserial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    updated_at bigint NOT NULL DEFAULT 0,
    updated_by bigserial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);