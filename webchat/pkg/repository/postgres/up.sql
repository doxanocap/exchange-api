CREATE TABLE chat_list (
    id SERIAL PRIMARY KEY,
    chat_uid VARCHAR(36),
    message_qty INT,
    started_at BIGINT,
    blocked BOOLEAN
);

CREATE INDEX chat_list_id
    ON chat_list (chat_uid);

CREATE TABLE chat_messages (
    chat_id INT REFERENCES chat_list(id),
    UNIQUE(chat_id),
    sender_id INT,
    sent_at BIGINT,
    message TEXT
);

CREATE TABLE chat_users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    phone_number TEXT,
    online BOOLEAN
);
