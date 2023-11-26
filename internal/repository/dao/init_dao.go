package dao

import "gorm.io/gorm"

func InitTables(database *gorm.DB) error {
	return database.AutoMigrate(&User{})
}
