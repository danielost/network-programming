-- +migrate Up

create table if not exists tictactoe
(
    id    uuid primary key,
    table jsonb,
);

-- +migrate Down

DROP TABLE tictactoe;