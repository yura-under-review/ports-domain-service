CREATE TABLE IF NOT EXISTS ports
(
    id          SERIAL,
    symbol      TEXT,
    name        TEXT,
    country     TEXT,
    province    TEXT,
    city        TEXT,
    alias       TEXT,
    regions     TEXT,
    timezones   TEXT,
    unlocks     TEXT,
    code        TEXT,
    lat         float8,
    lon         float8,

    constraint ports_pk primary key (id),
    constraint ports_symbol_unq unique (symbol)
);

CREATE INDEX IF NOT EXISTS ports_symbol_idx ON ports (symbol);