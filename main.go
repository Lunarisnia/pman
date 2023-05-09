/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/lunarisnia/pman/cmd"
	"github.com/lunarisnia/pman/data"
)

func main() {
	// TODO: FIX THIS ERROR:
	// D:/Personal/Workplace/golang-dojo/pman/data/password.go:18 file is not a database (26)
	defer data.EncryptFile()
	data.OpenDatabase()
	// data.MigrateDatabase()
	cmd.Execute()
}
