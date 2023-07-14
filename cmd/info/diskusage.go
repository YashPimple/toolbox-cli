/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

// diskusageCmd represents the diskusage command
var diskusageCmd = &cobra.Command{
	Use:   "diskusage",
	Short: "Prints disk usage of the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//logic
		defaultDirectory := "."

		usage := du.NewDiskUsage(defaultDirectory)

		fmt.Printf("Free disk space:%d in directory %s\n", usage.Free(), defaultDirectory)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskusageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskusageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
