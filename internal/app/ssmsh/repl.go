package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func startSSMRepl(cmd *SSMCommandLineTool) {
	// Initalize the SSM command line tool and start the REPL.
	fmt.Print(color.HiRedString("ssmsh ~ "))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if command == "exit" || command == "quit" {
			os.Exit(0)
		} else if command == "help" {
			fmt.Println("ssmsh - A tool for managing AWS Systems Manager")
			fmt.Println("Commands")
			fmt.Println("login - Login to your AWS account using shared credentials")
			fmt.Println("list  - List all instances managed by SSM in a given region")
			startSSMRepl(cmd)
		} else {
			cmd.commandType(command)
			startSSMRepl(cmd)
		}
	}
}
