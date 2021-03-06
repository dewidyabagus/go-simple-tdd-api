package user

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"go-simple-api/business/user"
)

type User struct {
	ID        uint         `gorm:"column:id;type:bigserial;primaryKey;not null"`
	UniqueID  string       `gorm:"column:unique_id;type:uuid;unique;index;default:gen_random_uuid();not null"`
	Username  string       `gorm:"column:username;type:varchar(100);not null"`
	Email     string       `gorm:"column:email;type:varchar(100);not null"`
	FirstName string       `gorm:"column:first_name;type:varchar(150);not null"`
	LastName  string       `gorm:"column:last_name;type:varchar(150);not null"`
	Password  string       `gorm:"column:password;type:varchar(128);not null"`
	Verify    bool         `gorm:"column:verify;type:boolean;default:false"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp with time zone;not null"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp with time zone;not null"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp with time zone;index"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) InsertNew(newUser *user.User) error {
	return r.db.Create(r.toInsertModel(newUser)).Error
}

func (r *Repository) GetByUsernameOrEmail(login string) (*user.User, error) {
	return nil, nil
}

func (r *Repository) UpdateByUniqueId(uniqueID string, updateUser *user.User) error {
	return nil
}

func (r *Repository) DeleteByUniqueId(uniqueID string) error {
	return nil
}

func (r *Repository) toInsertModel(u *user.User) *User {
	return &User{
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		Verify:    u.Verify,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
