package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var fileName string = "profile.json"

type Profile struct {
	Name string `json:"name"`
}

func initProfile(data Profile) error {
	file, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(fileName, file, 0644)

	return err
}

func isProfileExists() bool {
	exists := true
	_, err := os.Open(fileName)

	if err != nil {
		exists = false
	}

	return exists
}

func getProfile() Profile {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("[error] Unable to cropeneate file: ", err)
		panic("Unable to open file")
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var profile Profile
	json.Unmarshal(byteValue, &profile)

	return profile
}
