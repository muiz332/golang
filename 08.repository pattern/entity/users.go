package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id       int64
	Username string
	Email    string
	Gender   sql.NullString
	Password string
	CreateAt time.Time
}
