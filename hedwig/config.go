package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getConfigFile() *os.File {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	file, openErr := os.OpenFile(filepath.Join(homedir, ".hedwig"), os.O_RDWR, os.ModeAppend)
	if openErr == os.ErrNotExist || file == nil {
		file, err = os.Create(filepath.Join(homedir, ".hedwig"))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	return file
}

func readConfigFile() []string {
	configFile := getConfigFile()
	defer configFile.Close()
	var content bytes.Buffer
	_, err := io.Copy(&content, configFile)
	if err != nil {
		panic(err)
	}
	recipients := strings.Split(string(content.Bytes()), "\n")

	for i, recipient := range recipients {
		if recipient == "" {
			recipients = removeItem(recipients, i)
		}
	}

	return recipients
}

func removeItem(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func addRecipientToConfigFile(recipient string) {
	configFile := getConfigFile()
	defer configFile.Close()
	_, err := configFile.WriteString(recipient + "\n")
	if err != nil {
		panic(err)
	}
}