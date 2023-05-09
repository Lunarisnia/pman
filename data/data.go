package data

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/lunarisnia/pman/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenDatabase() error {
	config.CheckDatabaseFile()
	var err error

	db, err = gorm.Open(sqlite.Open(config.DBPATH), &gorm.Config{})
	fmt.Println(config.DBPATH)
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
