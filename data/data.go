package data

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/lunarisnia/pman/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func OpenDatabase() error {
	var databaseExisted = config.CheckDatabaseFile()
	var err error

	db, err = gorm.Open(sqlite.Open(config.DBPATH), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	if !databaseExisted {
		MigrateDatabase()
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

func errorHandler() {
	EncryptFile()
	log.Fatalln(`
==========================
Try running "pman init" then try again
==========================
	`)
}
