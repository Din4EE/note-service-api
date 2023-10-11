-- +goose Up
-- +goose StatementBegin
create table note (
    id bigserial primary key,
    title text,
    text text,
    author text,
    created_at timestamp not null default now(),
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table note;
-- +goose StatementEnd
