package main

import (
	"fmt"

	"github.com/fatih/color"
)

// SSMCommandLineTool determines the commands being run and uses the SSMClient
// to run the functions that correspond to the command input.
type SSMCommandLineTool struct {
	client        *SSMClient
	authenticated bool
}

// InitializeSSMCommandLineTool creates an unauthenticated command line object.
func InitializeSSMCommandLineTool() *SSMCommandLineTool {
	return &SSMCommandLineTool{
		client:        nil,
		authenticated: false,
	}
}

func (cmd *SSMCommandLineTool) commandType(command string) {
	switch command {
	case "list":
		if cmd.authenticated != false {
			cmd.listCommand()
		} else {
			fmt.Println("You must login to AWS to perform this operation.")
			return
		}
	case "login":
		if cmd.authenticated != true {
			cmd.loginCommand()
		} else {
			fmt.Println("You have already logged into your AWS account.")
			return
		}
	}
}

func (cmd *SSMCommandLineTool) loginCommand() {
	ssmClient, err := NewSSMClient()
	if err != nil {
		fmt.Println(color.RedString("An error occurred while authenticating..."))
		return
	}
	cmd.client = ssmClient
	cmd.authenticated = true
}

func (cmd *SSMCommandLineTool) listCommand() {
	cmd.client.ListManagedInstances()
	cmd.client.PrintManagedInstances()
}
