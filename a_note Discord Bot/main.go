package main

import (
	"fmt"

	"github.com/anurag/a_note_discord_bot/bot"
	"github.com/anurag/a_note_discord_bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
	bot.Start()

	<-make(chan struct{})
	return
}
