CREATE TABLE exchangers_keys (
    id SERIAL PRIMARY KEY,
    city VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE exchangers_currencies (
    exchanger_id INT REFERENCES exchangers (id),
    usd_buy FLOAT NOT NULL,
    usd_sell FLOAT NOT NULL,
    eur_buy FLOAT NOT NULL,
    eur_sell FLOAT NOT NULL,
    rub_buy FLOAT NOT NULL,
    rub_sell FLOAT NOT NULL,
    updated_time BIGINT NOT NULL
);

CREATE TABLE exchangers_info (
    exchanger_id INT REFERENCES exchangers (id),
    address TEXT NOT NULL,
    link TEXT NOT NULL,
    special_offer TEXT NOT NULL,
    phones TEXT NOT NULL
);

CREATE INDEX exchanger_id_index
    ON currency_data (exchanger_id);
