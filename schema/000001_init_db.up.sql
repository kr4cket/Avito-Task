CREATE TABLE IF NOT EXISTS segments
(
    id          serial          primary key,
    name        varchar(255)    not null unique,
    entirety    int
);

CREATE TABLE IF NOT EXISTS users
(
    id          serial          primary key,
    login       varchar(255)    not null unique,
    password    varchar(255)    not null
);

CREATE TABLE IF NOT EXISTS user_segments
(
    id              serial      primary key,
    user_id         int references users(id)    on delete cascade   not null,
    segment_name    varchar(255) references segments(name) on delete cascade   not null,
    expire    timestamp
);

CREATE TABLE IF NOT EXISTS operation_history
(
    id              serial                                          primary key,
    user_id         int references users(id)    on delete cascade   not null,
    segment_name    varchar(255) references segments(name) on delete cascade   not null,
    time            timestamp                                       not null default CURRENT_TIMESTAMP,
    is_delete       boolean                                         not null default false
);

CREATE UNIQUE INDEX ON user_segments (user_id, segment_name);