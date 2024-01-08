CREATE TABLE Accounts
(
    AccountID BIGSERIAL PRIMARY KEY,
    Name      text NOT NULL,
    Level     integer DEFAULT 0 NOT NULL,
    Email     text NOT NULL
);