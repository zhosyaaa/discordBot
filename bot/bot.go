package bot

import (
	"discordBot/bot/handlers"
	"discordBot/clients"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Bot represents a Discord bot with specific functionality.
type Bot struct {
	Token     string
	Session   *discordgo.Session
	Weather   *clients.WeatherClient
	Translate *clients.TranslateClient
	Commands  map[string]handlers.Command
}

// NewBot creates a new instance of Bot with the provided token, WeatherClient, and TranslateClient.
func NewBot(token string, weather *clients.WeatherClient, translate *clients.TranslateClient) *Bot {
	bot := &Bot{Token: token, Weather: weather, Translate: translate}
	bot.initCommands()
	return bot
}

// initCommands initializes the available commands for the bot.
func (b *Bot) initCommands() {
	b.Commands = map[string]handlers.Command{
		CommandHelp:      &handlers.HelpCommand{},
		CommandPoll:      &handlers.PollCommand{},
		CommandWeather:   &handlers.WeatherCommand{Weather: b.Weather},
		CommandTranslate: &handlers.TranslateCommand{Translate: b.Translate},
		CommandLanguage:  &handlers.LanguageCommand{},
	}
}

// Run starts the Discord bot and waits for an interrupt signal to gracefully shutdown.
func (b *Bot) Run() {
	// Create a new Discord session
	session, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		log.Printf("Error creating Discord session: %v", err)
		return
	}
	b.Session = session

	// Add a message handler for new messages
	session.AddHandler(b.newMessage)

	// Open the Discord session
	err = session.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer session.Close()

	// Print a message indicating the bot is running
	fmt.Println("Bot running....")

	// Wait for an interrupt signal to gracefully shutdown the bot
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

// newMessage is a handler for new messages received by the bot.
func (b *Bot) newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}
	// Trim leading and trailing whitespaces from the message content
	content := strings.TrimSpace(message.Content)

	// Check if the message starts with the bot's command prefix
	if strings.HasPrefix(content, CommandPrefix) {
		// Split the message into parts
		commandParts := strings.Fields(content)
		commandName := strings.TrimPrefix(commandParts[0], CommandPrefix)

		// Check if the command exists and execute it
		if cmd, exists := b.Commands[commandName]; exists {
			cmd.Execute(session, message, commandParts[1:])
		} else {
			// Send a message indicating that the command is unknown
			session.ChannelMessageSend(message.ChannelID, "Unknown handlers. Type '!help' for a list of commands.")
		}
	}
}
