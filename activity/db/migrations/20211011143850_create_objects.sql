-- +goose Up
-- +goose StatementBegin
set schema 'timesheet';

create table activities (
    id uuid primary key,
    owner uuid not null,
    type uuid not null,
    family uuid not null,
    date_from timestamp not null,
    date_to timestamp not null,
);

create table activity_types (
    id uuid primary key,
    family uuid not null,
    name varchar not null,
    description varchar not null,
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
set schema 'activity';

drop table activities;
drop table activity_types;

-- +goose StatementEnd