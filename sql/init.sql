CREATE TABLE users (
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null
);
CREATE TABLE sessions (
    id serial primary key,
    uuid varchar(64) not null unique,
    email varchar(255),
    user_id bigint unsigned references users(id),
    created_at timestamp not null
);
CREATE TABLE threads (
    id serial primary key,
    uuid varchar(64) not null unique,
    topic text,
    user_id bigint unsigned references users(id),
    created_at timestamp not null
);
CREATE TABLE posts (
    id serial primary key,
    uuid varchar(64) not null unique,
    body text,
    user_id bigint unsigned references users(id),
    thread_id bigint unsigned references threads(id),
    created_at timestamp not null
);