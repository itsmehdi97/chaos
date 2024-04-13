/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"chaos/uid"

	"github.com/spf13/cobra"
)

// uidCmd represents the uid command
var uidCmd = &cobra.Command{
	Use:   "uid",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		uid.Run()
	},
}

func init() {
	rootCmd.AddCommand(uidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
