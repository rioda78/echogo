package migrasi

import (
	"echogo/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.Migrator().DropColumn(&models.User{}, "Email")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.Response{})
}
