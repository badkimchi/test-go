CREATE TABLE Accounts
(
    Account_ID BIGSERIAL PRIMARY KEY,
    Name      text NOT NULL,
    Level     integer DEFAULT 0 NOT NULL,
    Email     text NOT NULL
);