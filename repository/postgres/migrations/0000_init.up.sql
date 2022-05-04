DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'network') THEN
            CREATE TYPE network AS ENUM ('ethereum', 'bsc', 'unknown');
        END IF;
    END
$$;


create table trades
(
    id         serial
        constraint trades_pk
            primary key,
    tx         varchar not null,
    tradehash  varchar not null,
    address_1  varchar,
    network_1  network,
    address_2  varchar,
    network_2  network default 'unknown',
    created_at timestamp default now()
);

create unique index trades_tradehash_uindex
    on trades (tradehash);

create table locks
(
    id          serial,
    tradehash   varchar
        constraint locks_trades_tradehash_fk
            references trades(tradehash)
            on update cascade on delete cascade,
    address     varchar,
    network     network,
    asset       varchar,
    amount      numeric,
    max_penalty numeric,
    deadline    bigint,
    created_at  timestamp default now()
);

create table engages
(
    id        serial,
    tradehash integer
        constraint locks_trades_tradehash_fk
            references trades
            on update cascade on delete cascade,
    address   varchar,
    network   network,
    signature varchar
);

create table claims
(
    id         serial,
    tradehash  integer
        constraint locks_trades_tradehash_fk
            references trades
            on update cascade on delete cascade,
    address    varchar,
    network    network,
    penalty    numeric,
    signatures varchar
);

CREATE TABLE IF NOT EXISTS blocks
(
    network network,
    number     integer                  NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

INSERT INTO blocks (network, number, updated_at)
VALUES ('ethereum', 10617362, now()),
       ('bsc', 19012549, now());


