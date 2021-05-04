
CREATE TABLE IF NOT EXISTS ports
(
    id    BIGINT       NOT NULL,
    symbol TEXT,
    name TEXT,
    city TEXT,
    province TEXT,
    alias TEXT,
    regions TEXT,
    timezones TEXT,
    unlocks TEXT,
    code TEXT,
    lat float8,
    lon float8,

    CONSTRAINT ports_pk PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS ports_symbol_idx ON ports (symbol);