-- +goose Up
-- +goose StatementBegin
alter table note add column email text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table note drop column email;
-- +goose StatementEnd
