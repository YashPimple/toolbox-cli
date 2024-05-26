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

var (
	clusterName  string
	region       string
	accessKey    string
	secretKey    string
	instanceType string
)

var CmdCreate = &cobra.Command{
	Use:   "create",
	Short: "Create an AWS EC2 instance",
	Run:   createEC2Instance,
}

func readPassword() (string, error) {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return string(bytePassword), nil
}

func createEC2Instance(cmd *cobra.Command, args []string) {
	godotenv.Load()

	fmt.Println("Enter Name of your Cluster:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading the input:", err)
		return
	}
	clusterName = strings.TrimSpace(input)
	fmt.Printf("Cluster Name: %s\n", clusterName)

	fmt.Println("Enter AWS Region(us-east-1):")
	awsRegion, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in Reading th input:", err)
		return
	}
	region = strings.TrimSpace(awsRegion)

	fmt.Println("Enter AWS EC2 Instance Type (e.g., t2.micro):")
	output, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading the input:", err)
		return
	}
	instanceType = strings.TrimSpace(output)

	fmt.Println("Enter your Access Key:")
	accessKeyInput, err := readPassword()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	accessKey = strings.TrimSpace(accessKeyInput)

	// Get secret key
	fmt.Println("Enter your Secret Key:")
	secretKeyInput, err := readPassword()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	secretKey = strings.TrimSpace(secretKeyInput)

	// Initialize AWS session
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		fmt.Println("Error initializing AWS session:", err)
		os.Exit(1)
	}

	// Create an EC2 service client
	ec2Client := ec2.New(awsSession)

	// Specifying  details
	runInput := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-04b70fa74e45c3917"),
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	result, err := ec2Client.RunInstances(runInput)

	if err != nil {
		fmt.Println("Error creating EC2 instance:", err)
		os.Exit(1)
	}

	instanceID := *result.Instances[0].InstanceId
	fmt.Println("Initalizing EC2 Instance")
	fmt.Printf("Created EC2 instance with ID: %s\n", instanceID)
}

func init() {

}
