package migration

import (
	"gorm.io/gorm"

	"go-simple-api/modules/main/user"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(user.User{})
}
