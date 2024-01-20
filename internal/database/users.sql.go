// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (
        email, name, password, profile_picture, role
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    id
`

type CreateUserParams struct {
	Email          string
	Name           string
	Password       string
	ProfilePicture sql.NullString
	Role           UserRole
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Name,
		arg.Password,
		arg.ProfilePicture,
		arg.Role,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const doesUserExists = `-- name: DoesUserExists :one
SELECT EXISTS (
        SELECT id, role, name, profile_picture, email, password, created_at, updated_at
        FROM users
        WHERE
            email = $1
    )
`

func (q *Queries) DoesUserExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, doesUserExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, role, name, profile_picture, email, password, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Name,
		&i.ProfilePicture,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
