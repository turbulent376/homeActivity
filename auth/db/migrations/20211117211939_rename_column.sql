-- +goose Up
-- +goose StatementBegin
set schema 'auth';

alter table sessions rename column device_id to device_name; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
set schema 'auth';

alter table sessions rename column device_name to device_id;
-- +goose StatementEnd
