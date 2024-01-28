/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/spf13/cobra"
)

// ulidCmd represents the ulid command
var ulidCmd = &cobra.Command{
	Use:   "ulid",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for range make([]int, n) {
			fmt.Println(ulid.Make())
		}
	},
}

func init() {
	rootCmd.AddCommand(ulidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ulidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ulidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
