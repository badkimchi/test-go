-- name: GetAccount :one
SELECT *
FROM Account
WHERE AccountID = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO Account (Name, Password, Level, Email)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateAccount :exec
UPDATE Account
SET Name     = $2,
    Password = $3,
    Level    = $4,
    Email    = $5
WHERE AccountID = $1;

-- name: DeleteAccount :exec
DELETE
FROM Account
WHERE AccountID = $1;