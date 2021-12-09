-- +goose Up
-- +goose StatementBegin
set schema 'auth';

alter table sessions add created_at timestamp;
alter table sessions add client_version varchar; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
set schema 'auth';

alter table sessions drop column created_at;
alter table sessions drop column client_version;
-- +goose StatementEnd
