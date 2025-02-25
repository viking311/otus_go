-- +goose Up
-- +goose StatementBegin
CREATE TABLE events(
                       id          uuid default gen_random_uuid() not null primary key,
                       title       varchar(255)                   not null,
                       description text,
                       event_time    timestamp                      not null,
                       duration    varchar(255),
                       remind_time integer,
                       user_id     integer                        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events
-- +goose StatementEnd