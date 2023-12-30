CREATE TABLE Account
(
    AccountID BIGSERIAL PRIMARY KEY,
    Name      text,
    Password  text NOT NULL,
    Level     integer DEFAULT 0,
    Email     text
);