/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/lunarisnia/pman/data"
	"github.com/spf13/cobra"
)

var serviceName string
var passwordLength int

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a password",
	Long: `Generate a password for your need.
	
	Usage:
	pman generate --name facebook --length 16`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.DecryptFile()
		generatedPassword := ""
		for i := 0; i < passwordLength; i++ {
			generatedPassword += randomChar()
		}
		data.InsertPassword(serviceName, generatedPassword)

		fmt.Fprintf(cmd.OutOrStdout(), "Generated Password: %v\n", generatedPassword)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringVarP(&serviceName, "name", "n", "", "Service name your password is tied to (required)")
	generateCmd.MarkFlagRequired("name")

	generateCmd.Flags().IntVarP(&passwordLength, "length", "l", 16, "Your password length default to 16")
}

func randomChar() string {
	r := rune(randomNumberInRange(33, 126))
	return string(r)
}

func randomNumberInRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
