CREATE TABLE exchangers_keys (
    id SERIAL PRIMARY KEY,
    city VARCHAR(100),
    name VARCHAR(100)
);

CREATE TABLE exchangers_currencies (
    exchanger_id INT REFERENCES exchangers_keys(id),
    upload_time BIGINT,
    usd_buy FLOAT,
    usd_sell FLOAT,
    eur_buy FLOAT,
    eur_sell FLOAT,
    rub_buy FLOAT,
    rub_sell FLOAT
);

CREATE TABLE exchangers_info (
    exchanger_id INT REFERENCES exchangers_keys(id),
    address TEXT,
    wholesale TEXT,
    updated_time BIGINT,
    phone_numbers TEXT
);

CREATE INDEX exchanger_keys_index
    ON exchangers_currencies (exchanger_id);
    