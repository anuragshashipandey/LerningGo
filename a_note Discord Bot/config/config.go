package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"` // this last string term helps us to encode it in the desired format i.e camelcase as thats what ia  followed in JSON
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading the config file")

	file, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Error while Reading Creditianls File")
		return err
	}

	// fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
