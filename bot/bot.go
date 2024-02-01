package bot

import (
	"discordBot/clients"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

type Bot struct {
	Token           string
	Session         *discordgo.Session
	Weather         *clients.WeatherClient
	ReminderManager *ReminderManager
}

func NewBot(token string, weather *clients.WeatherClient) *Bot {
	return &Bot{Token: token, Weather: weather, ReminderManager: &ReminderManager{}}
}

func (b *Bot) Run() {
	session, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		log.Printf("Error creating Discord session: %v", err)
		return
	}
	b.Session = session

	session.AddHandler(b.newMessage)

	err = session.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer session.Close()

	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
