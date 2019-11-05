package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func createSSMSHDirectory() error {
	err := os.Mkdir("./.ssmsh/", 0777)
	if err != nil {
		return err
	}
	return nil
}

func createSSMSHConfigurationFile() error {
	configuration := []byte("{\n\t\"region\" : \"\",\n\t\"profile\":\"\"\n}")
	err := ioutil.WriteFile("./.ssmsh/config.json", configuration, 0660)
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

func extractSSMSessionPlugin() error {
	err := exec.Command(
		"unzip",
		".ssmsh/sessionmanager-bundle.zip",
		"-d",
		".ssmsh/",
	).Run()
	if err != nil {
		return err
	}
	return nil
}

func installSSMSessionPlugin() error {
	if runtime.GOOS == "darwin" {
		err := exec.Command(
			"cp",
			"-rf",
			".ssmsh/sessionmanager-bundle/bin/",
			"/usr/local/",
		).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
