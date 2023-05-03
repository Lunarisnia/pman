package data

import (
	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	ServiceName string
	Value       string
}

func InsertPassword(serviceName string, value string) {
	password := Password{ServiceName: serviceName, Value: value}
	if err := db.Create(&password).Error; err != nil {
		EncryptFile()
		remindInit()
	}
}

func ReadAllPasswords() []Password {
	passwords := []Password{}
	if err := db.Find(&passwords).Error; err != nil {
		EncryptFile()
		remindInit()
	}
	return passwords
}

func ReadOnePassword(passwordID string) Password {
	p := Password{}
	if err := db.Where("id = ?", passwordID).First(&p).Error; err != nil {
		EncryptFile()
		remindInit()
	}

	return p
}
