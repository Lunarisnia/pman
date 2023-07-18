package config

import (
	"fmt"
	"os"
)

var DBPATH = ConfigFilePath() + "/pman-vault.pman"

func ConfigFilePath() string {
	configDir, _ := os.UserConfigDir()
	return configDir + "/pman"
}

func CheckDatabaseFile() bool {
	filepath := DBPATH
	if _, err := os.ReadDir(ConfigFilePath()); err != nil {
		os.Mkdir(ConfigFilePath(), 0777)
	}
	_, err := os.ReadFile(filepath)
	fmt.Println(err, "=====")

	return err == nil
}
