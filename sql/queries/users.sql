-- name: CreateUser :one
INSERT INTO
    users (
        email, name, password, profile_picture, role
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: DoesUserExists :one
SELECT EXISTS (
        SELECT *
        FROM users
        WHERE
            email = $1
    );