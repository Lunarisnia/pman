/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lunarisnia/pman/data"
	"github.com/spf13/cobra"
)

var passwordID string

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Reveal your password",
	Long:  `Reveal your password`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.DecryptFile()
		password := data.ReadOnePassword(passwordID)
		data.EncryptFile()

		fmt.Printf("%v: %v", password.ServiceName, password.Value)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&passwordID, "id", "i", "", "Your password ID")
	showCmd.MarkFlagRequired("id")
}
