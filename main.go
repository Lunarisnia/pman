/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/lunarisnia/pman/cmd"
	"github.com/lunarisnia/pman/data"
)

func main() {
	defer data.EncryptFile()
	data.DecryptFile()
	data.OpenDatabase()
	cmd.Execute()
}
