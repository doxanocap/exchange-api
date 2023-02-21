CREATE TABLE exchangers_keys (
     id SERIAL PRIMARY KEY,
     city VARCHAR(100),
     name VARCHAR(100)
);

CREATE TABLE exchangers_currencies (
    exchanger_id INT PRIMARY KEY,
    usd_buy FLOAT,
    usd_sell FLOAT,
    eur_buy FLOAT,
    eur_sell FLOAT,
    rub_buy FLOAT,
    rub_sell FLOAT,
    updated_time BIGINT
);

CREATE TABLE exchangers_info (
     exchanger_id INT PRIMARY KEY,
     address TEXT,
     link TEXT,
     special_offer TEXT,
     phone_numbers TEXT
);

CREATE INDEX exchanger_keys_index
    ON exchangers_currencies (exchanger_id);
    