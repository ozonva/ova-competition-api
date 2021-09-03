-- +goose Up
create table competitions
(
    id         bigserial               not null unique,
    name       varchar(250) default '' not null,
    start_time timestamp               not null,
    status     integer                 not null
);

-- +goose Down
DROP TABLE competitions;