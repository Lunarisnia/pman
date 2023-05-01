/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lunarisnia/pman/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiate the database for the program",
	Long:  `Initiate the database for the program`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.MigrateDatabase()
		fmt.Fprintln(cmd.OutOrStdout(), "Database Initiated")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
