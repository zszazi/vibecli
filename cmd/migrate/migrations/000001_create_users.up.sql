CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY, 
    email citext UNIQUE NOT NULL, -- guy@mail.com is same as GuY@mAil.com - to make emails stored in lowercase
    username varchar(255) UNIQUE NOT NULL, 
    password bytea NOT NULL, 
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
)

