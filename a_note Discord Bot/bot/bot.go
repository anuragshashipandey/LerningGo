package bot

import (
	"fmt"

	"github.com/anurag/a_note_discord_bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println("Bot Oppening Error", err.Error())
		return
	}
	fmt.Println("Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	println("Msg from", m.Author.Username, m.Author.GlobalName)
	if m.Author.ID == BotId {
		return
	}
	if m.Content != "" {
		var reciever string = m.Author.Username
		fmt.Println(reciever)
		if m.Author.Username == "i_am_pandey" {
			_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hi! Owner....%s.... How are you", reciever))
			if err != nil {
				fmt.Println("Cant Send msg", err.Error())
			}
		} else {
			_, err := s.ChannelMessageSend(m.ChannelID, "Hi!")

			if err != nil {
				fmt.Println("Message Sending Error", err.Error())
			}
		}
	}
}
