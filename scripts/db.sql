create database balance_service;
create schema balance_service;

create table balance_service.users
(
    id      serial primary key,
    user_id varchar(50) not null,
    balance int not null default 0
);

-- create table balance_service.transactions
-- (
--     id      serial primary key,
--     user_id varchar(50) not null,
--     type varchar(20)
-- );