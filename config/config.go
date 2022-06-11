package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

//configstruct are your own data types in golang
type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json") //read the file contents

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, config) //unmarshall the file and bring the values into the config

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//assign the config
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}
