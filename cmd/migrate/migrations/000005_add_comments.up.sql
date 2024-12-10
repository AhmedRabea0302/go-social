CREATE TABLE IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    user_id bigserial NOT NULL,
    post_id bigserial Not NULL,
    content TEXT NOT NULL,
    created_at timestamp(0) with time zone Not NULL DEFAULT NOW()
)