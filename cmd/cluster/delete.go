/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createClusterCmd represents the createCluster command
var Delete = &cobra.Command{
	Use:   "delete",
	Short: "Use to delete a cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster deleted")
	},
}

func init() {

}
