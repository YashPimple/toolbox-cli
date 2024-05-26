package cluster

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var CmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete an AWS EC2 instance",
	Run:   deleteEC2Instance,
}

func getPassword() (string, error) {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a new line after the input
	return string(bytePassword), nil
}

func deleteEC2Instance(cmd *cobra.Command, args []string) {
	godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// 	return
	// }

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter AWS Region (e.g., us-east-1):")
	regionInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	region := strings.TrimSpace(regionInput)

	fmt.Println("Enter EC2 Instance ID to delete:")
	instanceIDInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	instanceID := strings.TrimSpace(instanceIDInput)

	// Get access key
	fmt.Println("Enter your Access Key:")
	accessKeyInput, err := getPassword()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	accessKey := strings.TrimSpace(accessKeyInput)

	// Get secret key
	fmt.Println("Enter your Secret Key:")
	secretKeyInput, err := getPassword()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	secretKey := strings.TrimSpace(secretKeyInput)

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}

	// Create EC2 service client
	svc := ec2.New(sess)

	// Terminate the EC2 instance
	terminateInput := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	_, err = svc.TerminateInstances(terminateInput)
	if err != nil {
		fmt.Printf("Error terminating EC2 instance: %v\n", err)
		return
	}

	fmt.Println("Successfully requested termination of EC2 instance:", instanceID)
}

func init() {

}
