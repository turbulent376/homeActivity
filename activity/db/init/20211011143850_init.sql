-- +goose Up
CREATE ROLE timesheet LOGIN PASSWORD 'activity' NOINHERIT CREATEDB;
CREATE SCHEMA activity AUTHORIZATION activity;
GRANT USAGE ON SCHEMA activity TO PUBLIC;

-- +goose Down
DROP SCHEMA activity;
DROP ROLE activity;
