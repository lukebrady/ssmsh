package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMClient holds an SSM Client object with associated to an AWS session.
type SSMClient struct {
	client           *ssm.SSM
	managedInstances []*string
	configuration    *SSMClientConfiguration
}

// SSMClientConfiguration holds the SSM configuration from .ssmsh/config.json
type SSMClientConfiguration struct {
	Region  string `json:"region"`
	Profile string `json:"profile"`
}

// NewSSMClient creates a new AWS session and returns the SSM client object.
func NewSSMClient() (*SSMClient, error) {
	// Read in the SSM configuration.
	configObject := &SSMClientConfiguration{}
	configFile, err := ioutil.ReadFile(".ssmsh/config.json")
	if err != nil {
		fmt.Println("Have you initalized SSMSH? Run init to create the .ssmsh directory.")
		return nil, err
	}
	json.Unmarshal(configFile, configObject)
	// Create a new session with static credential values.
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		Profile:                 configObject.Profile,
	})
	if err != nil {
		return nil, err
	}
	return &SSMClient{
		client:        ssm.New(sess),
		configuration: configObject,
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

// StartSSMSession creates a new SSM session and enters the remote instance.
func (s *SSMClient) StartSSMSession(instanceID string) error {
	err := exec.Command(
		"aws",
		"ssm",
		"start-session",
		"--target",
		instanceID,
	).Run()
	if err != nil {
		return err
	}
	return nil
}

// PrintManagedInstances prints the list of managed instances.
func (s *SSMClient) PrintManagedInstances() {
	for _, instance := range s.managedInstances {
		fmt.Println(*instance)
	}
}
