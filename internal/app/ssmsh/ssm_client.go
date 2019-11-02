package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMClient holds an SSM Client object with associated to an AWS session.
type SSMClient struct {
	client           *ssm.SSM
	managedInstances []*string
}

// NewSSMClient creates a new AWS session and returns the SSM client object.
func NewSSMClient() (*SSMClient, error) {
	// Create a new session with static credential values.
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	})
	if err != nil {
		return nil, err
	}
	return &SSMClient{
		client: ssm.New(sess),
	}, nil
}

// ListManagedInstances lists out all instances within a region and caches the values.
func (s *SSMClient) ListManagedInstances() {
	params := &ssm.DescribeInstanceInformationInput{}
	err := s.client.DescribeInstanceInformationPages(params, func(page *ssm.DescribeInstanceInformationOutput, last bool) bool {
		for _, instance := range page.InstanceInformationList {
			s.managedInstances = append(s.managedInstances, instance.InstanceId)
		}
		if page.NextToken == nil {
			return false
		}
		return true
	})
	if err != nil {
		log.Println(err)
	}
}

// PrintManagedInstances prints the list of managed instances.
func (s *SSMClient) PrintManagedInstances() {
	for _, instance := range s.managedInstances {
		fmt.Println(*instance)
	}
}
