create database balance_service;

create schema balance_service;

create table balance_service.users
(
    id      serial primary key,
    user_id int not null,
    balance int not null default 0,
);
