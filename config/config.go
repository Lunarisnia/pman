package config

import (
	"os"
	"runtime"
)

func currentOS() string {
	return runtime.GOOS
}

func ConfigFilePath() string {
	switch currentOS() {
	case "windows":
		return "%PROGRAMDATA%/pman"
	default:
		return "etc/pman"
	}
}

func CheckDatabaseFile() bool {
	filepath := ConfigFilePath() + "/pman-vault.pman"
	_, err := os.ReadFile(filepath)

	return err == nil
}
