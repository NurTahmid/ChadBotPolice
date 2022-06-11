package main

import (
	"fmt"

	"github.com/NurTahmid/ChadBotPolice/bot"
	"github.com/NurTahmid/ChadBotPolice/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
