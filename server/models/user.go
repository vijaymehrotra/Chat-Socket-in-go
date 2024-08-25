package models

import (
	// "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`
}

func AutoMigrate(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}