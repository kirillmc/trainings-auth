-- +goose Up
create table users
(
    id            serial primary key,
    name          varchar(50)  not null,
    surname       varchar(50)  not null,
    email         varchar(255) not null unique,
    avatar        varchar(255) not null,
    login         varchar(50)  not null unique,
    password_hash varchar(255) not null,
    role          integer      not null,
    locked        boolean      not null default false
);

create table roles_to_endpoints
(
    id       serial primary key,
    endpoint text    not null,
    role     integer not null
);


insert into roles_to_endpoints (endpoint, role)
values ('/training_v1.TrainingV1/CreateTrainingProgram', 3),
       ('/user_v1.UserV1/Get', 3),
       ('/user_v1.UserV1/Delete', 3),
       ('/user_v1.UserV1/UpdatePassword', 3),
       ('/user_v1.UserV1/UpdateRole', 3),
       ('/user_v1.UserV1/UnlockUser', 3),
       ('/user_v1.UserV1/UpdateRole', 2),
       ('/user_v1.UserV1/UnlockUser', 2),
       ('/user_v1.UserV1/LockUser', 3),
       ('/user_v1.UserV1/Update', 3);
-- +goose Down
drop table users;
drop table roles_to_endpoints;

