package data

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenDatabase() error {
	var err error

	// TODO: add encryption for the database file
	db, err = gorm.Open(sqlite.Open("pman-vault.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func MigrateDatabase() {
	err := db.AutoMigrate(&Password{})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	log.Println("Database initiated")
}

func remindInit() {
	log.Fatalln(`
==========================
Have you ran "pman init"?
==========================
	`)
}
