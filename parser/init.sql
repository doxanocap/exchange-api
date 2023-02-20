CREATE TABLE currency_data (
    exchanger_id INT REFERENCES exchangers (id),
    usd_buy INT NOT NULL,
    usd_sell INT NOT NULL,
    eur_buy INT NOT NULL,
    eur_sell INT NOT NULL,
    rub_buy INT NOT NULL,
    rub_sell INT NOT NULL,
    timestamp TIMESTAMP DEFAULT NOW()
);

CREATE TABLE exchangers (
    id SERIAL PRIMARY KEY,
    city VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE exchangers (
    id SERIAL PRIMARY KEY,
    city VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    link TEXT NOT NULL,
    address TEXT NOT NULL,
    sale TEXT NOT NULL,
    phones TEXT NOT NULL
);

CREATE INDEX exchanger_id_index
ON currency_data (exchanger_id);

CREATE INDEX exchangers_city_index
ON exchangers (city);