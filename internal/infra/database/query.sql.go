// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const createUserExternal = `-- name: CreateUserExternal :one
INSERT INTO users (
  name,
  email,
  provider
) VALUES ($1, $2, $3) RETURNING id, name, email, provider, password, created_at, updated_at
`

type CreateUserExternalParams struct {
	Name     string
	Email    string
	Provider sql.NullString
}

func (q *Queries) CreateUserExternal(ctx context.Context, arg CreateUserExternalParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUserExternal, arg.Name, arg.Email, arg.Provider)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Provider,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserLocal = `-- name: CreateUserLocal :one
INSERT INTO users (
  name,
  email,
  password
) VALUES ($1, $2, $3) RETURNING id, name, email, provider, password, created_at, updated_at
`

type CreateUserLocalParams struct {
	Name     string
	Email    string
	Password sql.NullString
}

func (q *Queries) CreateUserLocal(ctx context.Context, arg CreateUserLocalParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUserLocal, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Provider,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, provider, password, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Provider,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, provider, password, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Provider,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET name = $2, password = $3, provider = $4 WHERE id = $1 RETURNING id, name, email, provider, password, created_at, updated_at
`

type UpdateUserParams struct {
	ID       int64
	Name     string
	Password sql.NullString
	Provider sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Password,
		arg.Provider,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Provider,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
