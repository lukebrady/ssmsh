package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMClient holds an SSM Client object with associated to an AWS session.
type SSMClient struct {
	client *ssm.SSM
}

// NewSSMClient creates a new AWS session and returns the SSM client object.
func NewSSMClient(profileName string) *SSMClient {
	session, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		panic(err)
	}
	return &SSMClient{
		client: ssm.New(session),
	}
}

// ListManagedInstances lists out all instances within a region and caches the values.
func (s *SSMClient) ListManagedInstances() {
	managedInstances, err := s.client.DescribeInstanceInformation(&ssm.DescribeInstanceInformationInput{})
	if err != nil {
		panic(err)
	}
	fmt.Println(managedInstances.String())
}
