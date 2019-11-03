package main

import (
	"os"
	"os/exec"
)

func createSSMSHDirectory() error {
	err := os.Mkdir("./.ssmsh/", 0777)
	if err != nil {
		return err
	}
	return nil
}

func downloadSSMSessionPlugin() error {
	// Curl the SSM session manager bundle and extact.
	err := exec.Command(
		"curl",
		"https://s3.amazonaws.com/session-manager-downloads/plugin/latest/mac/sessionmanager-bundle.zip",
		"-o",
		"./.ssmsh/sessionmanager-bundle.zip",
	).Run()
	if err != nil {
		return err
	}
	return nil
}
