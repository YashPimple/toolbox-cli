/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cluster

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your AWS Instance",
	Run:   deleteEC2Instance,
}

func deleteEC2Instance(cmd *cobra.Command, args []string) {

	var instanceId string
	var region string
	var accessKey string
	var secretKey string
	fmt.Println("Enter the ID of the EC2 instance to delete : ")
	fmt.Scan(&instanceId)

	fmt.Println("Enter AWS Region(us-east-1):")
	fmt.Scan(&region)

	fmt.Println("Enter your Access Key:")
	fmt.Scan(&accessKey)

	fmt.Println("Enter Secret Key:")
	fmt.Scan(&secretKey)

	// Initializing AWS session
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		fmt.Println(err)
	}

	// Creating an EC2 service client
	ec2Client := ec2.New(awsSession)

	terimateInput := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(instanceId)},
	}

	// Terminating
	_, err = ec2Client.TerminateInstances(terimateInput)

	if err != nil {
		fmt.Println("Error in deleting EC2 instance: ", err)
		os.Exit(1)
	}

	fmt.Printf("Terminated EC2 instance with ID: %s\n", instanceId)

}

func init() {
}
