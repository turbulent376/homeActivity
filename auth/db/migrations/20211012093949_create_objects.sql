-- +goose Up
set schema 'auth';

create table users
(
  id uuid primary key,
  name varchar,
  surname varchar,
  avatar varchar,
  firebase_uuid varchar,
  kundelik_id varchar,
  email varchar not null,
  password varchar not null,
  country_code varchar,
  created_at timestamp not null,
  updated_at timestamp not null,
  deleted_at timestamp null
);

-- +goose Down
drop table users;

