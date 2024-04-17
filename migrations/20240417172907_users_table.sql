-- +goose Up
create table users
(
    id         serial primary key,
    name   varchar(50)      not null,
    surname   varchar(50)      not null,
    email      varchar(255)      not null unique,
    avatar      varchar(255)      not null,
    login   varchar(50)      not null unique,
    password_hash   varchar(255)      not null,
    locked boolean not null default false
);

create table admins
(
    id       serial primary key,
    user_id int    not null unique,
    foreign key (user_id) references users(id) on delete cascade
);

create table moders
(
    id       serial primary key,
    user_id int    not null unique,
    foreign key (user_id) references users(id) on delete cascade
);

create table roles_to_endpoints
(
    id       serial primary key,
    endpoint text    not null,
    role     integer not null
);

insert into roles_to_endpoints (endpoint, role)
values ('/chat_v1.ChatV1/Create', 2),
       ('/user_v1.UserV1/Get', 2),
       ('/user_v1.UserV1/Delete', 2),
       ('/user_v1.UserV1/Update', 2);
-- +goose Down
drop table users;
drop table admins;
drop table moders;
drop table roles_to_endpoints;

