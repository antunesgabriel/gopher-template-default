// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Provider  sql.NullString
	Password  sql.NullString
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
