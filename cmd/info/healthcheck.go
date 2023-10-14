package info

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var domain string

func healthcheck(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v\n", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}

// healthcheckCmd represents the healthcheck command
var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Command that checks the given domain is down or not",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//logic
		status := healthcheck(domain, "80")
		fmt.Println(status)
	},
}

func init() {
	healthcheckCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain to check health")

	if err := healthcheckCmd.MarkFlagRequired("domain"); err != nil {
		fmt.Println(err)
	}
}
