-- +goose Up
-- +goose StatementBegin
set schema 'auth';

create table sessions
(
  id uuid primary key,
  user_id uuid not null,
  refresh_token text,
  device_id varchar,
  fcm_token varchar

);

create index sessions_user_id_index on sessions(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
set schema 'auth';

drop table sessions;
-- +goose StatementEnd
