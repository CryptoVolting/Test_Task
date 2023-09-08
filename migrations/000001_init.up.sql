CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS operators
(
    id              uuid            not null unique default uuid_generate_v4(),
    name            varchar(255)    not null,
    surname         varchar(255)    not null,
    town            varchar(255)    not null,
    telephone       bigint          not null unique check (telephone >= 80000000000 and telephone <= 89999999999),
    email           varchar(255)    not null unique,
    password        varchar(255)    not null
);

CREATE TABLE IF NOT EXISTS projects
(
    id          uuid            not null unique default uuid_generate_v4(),
    name        varchar(255)    not null unique,
    typeProject varchar(255)    not null check (typeProject = 'входящий' OR typeProject = 'исходящий' OR typeProject = 'автоинформатор')
);

CREATE TABLE IF NOT EXISTS project_operators_list
(
    id              serial          not null,
    operators_id    uuid            not null references operators (id) on delete cascade,
    project_id      uuid            not null references projects (id) on delete cascade
);