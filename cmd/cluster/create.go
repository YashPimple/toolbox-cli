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

func createEC2Instance(cmd *cobra.Command, args []string) {

	fmt.Println("Enter Name of your Cluster:")
	fmt.Scan(&clusterName)

	fmt.Println("Enter AWS Region(us-east-1):")
	fmt.Scan(&region)

	fmt.Println("Enter AWS EC2 Instance Type (e.g., t2.micro):")
	fmt.Scan(&instanceType)

	fmt.Println("Enter your Access Key:")
	fmt.Scan(&accessKey)

	fmt.Println("Enter Secret Key:")
	fmt.Scan(&secretKey)

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
		ImageId:      aws.String("ami-0a21e01face015dd9"),
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
