package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
	case "init":
		cmd.initCommand()
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
	case "session":
		if cmd.authenticated != false {
			fmt.Print("\nEnter the instance id: ")
			scanner := bufio.NewScanner(os.Stdin)
			for {
				instanceID := scanner.Text()
				cmd.client.StartSSMSession(instanceID)
			}
		} else {
			fmt.Println("You must login to AWS to perform this operation.")
			return
		}
	default:
		fmt.Printf("%s is not an ssmsh command.\n", command)
	}
}

func (cmd *SSMCommandLineTool) initCommand() {
	err := createSSMSHDirectory()
	if err != nil {
		log.Println(err)
		return
	}
	err = createSSMSHConfigurationFile()
	if err != nil {
		log.Println(err)
		return
	}
	err = downloadSSMSessionPlugin()
	if err != nil {
		log.Println(err)
		return
	}
	err = extractSSMSessionPlugin()
	if err != nil {
		return
	}
	err = installSSMSessionPlugin()
	if err != nil {
		return
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

func (cmd *SSMCommandLineTool) startSessionCommand(instanceID string) {
	cmd.client.StartSSMSession(instanceID)
}

func (cmd *SSMCommandLineTool) listCommand() {
	cmd.client.ListManagedInstances()
	cmd.client.PrintManagedInstances()
}
