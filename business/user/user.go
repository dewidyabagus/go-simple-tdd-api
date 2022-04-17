package user

import (
	"time"
)

type User struct {
	ID        uint
	UniqueID  string
	Username  string
	Email     string
	FirstName string
	LastName  string
	Password  string
	Verify    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
