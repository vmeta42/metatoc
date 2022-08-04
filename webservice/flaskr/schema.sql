DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS transactions;

CREATE TABLE users (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       wallet_address TEXT UNIQUE NOT NULL,
       private_key TEXT UNIQUE NOT NULL
);

CREATE TABLE transactions (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       transaction_type TEXT,
       from_address TEXT,
       to_address TEXT,
       token_name TEXT,
       file_name TEXT
);
