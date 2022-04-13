DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'network') THEN
            CREATE TYPE network AS ENUM ('ethereum', 'bsc');
        END IF;
    END
$$;


create table trades
(
    id               serial
        constraint trades_pk
            primary key,
    tradehash        varchar not null,
    address_1        varchar,
    network_1        network,
    asset_1          varchar,
    amount_1         bigint,
    address_2        varchar,
    lock_tradehash_1 varchar,
    withdraw_1       varchar,
    network_2        network,
    asset_2          varchar,
    amount_2         bigint,
    withdraw_2       varchar,
    lock_tradehash_2 varchar,
    created_at       timestamp default now(),
    updated_at       timestamp
);

create unique index trades_tradehash_uindex
    on trades (tradehash);

