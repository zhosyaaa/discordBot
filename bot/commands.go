package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const (
	CommandPrefix    = "!"
	CommandHelp      = "help"
	CommandBye       = "bye"
	CommandPoll      = "poll"
	CommandWeather   = "weather"
	CommandTranslate = "translate"
	CommandRemind    = "remind"
	CommandGame      = "game"
)

func (b *Bot) newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Full Message Content:", message.Content)
	content := strings.TrimSpace(message.Content)
	if strings.HasPrefix(content, CommandPrefix) {
		_ = strings.TrimPrefix(content, CommandPrefix)
		command := strings.Fields(message.Content)[0:]

		fmt.Println(command[0])
		switch strings.TrimPrefix(command[0], CommandPrefix) {
		case CommandHelp:
			b.sendHelpMessage(message.ChannelID)
		case CommandBye:
			session.ChannelMessageSend(message.ChannelID, "Good Bye👋")
		case CommandPoll:
			b.handlePollCommand(session, message)
		case CommandWeather:
			b.handleWeatherCommand(session, message)
		case CommandTranslate:
			b.handleTranslateCommand(session, message)
		case CommandRemind:
			b.handleRemindCommand(session, message)
		case CommandGame:
			b.handleGameCommand(session, message)
		default:
			session.ChannelMessageSend(message.ChannelID, "Unknown command. Type '!help' for a list of commands.")
		}
	}
}

func (b *Bot) sendHelpMessage(channelID string) {
	helpMessage := "Available commands:\n" +
		"!help - Show this help message\n" +
		"!bye - Say goodbye\n" +
		"!poll [question] [option1] [option2] ... - Create a poll\n" +
		"!weather [location] - Get current weather information\n" +
		"!translate [language code] [text] - Translate text to the specified language\n" +
		"!remind [time] [message] - Set a reminder\n" +
		"!game - Play a text-based game\n"
	b.Session.ChannelMessageSend(channelID, helpMessage)
}

func (b *Bot) handlePollCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Split the content by spaces
	words := strings.Fields(message.Content)

	// Check if there are enough words in the message
	if len(words) < 4 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !poll [question] [option1] [option2] ...")
		return
	}

	// Extract the question and options
	question := words[1]
	optionList := words[2:]

	// Check if there are at least two options
	if len(optionList) < 2 {
		session.ChannelMessageSend(message.ChannelID, "You need at least two options to create a poll.")
		return
	}

	pollMessage := "Poll: " + question + "\n"
	for i, option := range optionList {
		pollMessage += fmt.Sprintf("%d. %s\n", i+1, option)
	}

	pollMessageSent, err := session.ChannelMessageSend(message.ChannelID, pollMessage)
	if err != nil {
		log.Printf("Error sending poll message: %v", err)
		return
	}

	for i := range optionList {
		err := session.MessageReactionAdd(message.ChannelID, pollMessageSent.ID, emojiNumber[i])
		if err != nil {
			log.Printf("Error adding reaction: %v", err)
			return
		}
	}
}

var emojiNumber = []string{
	"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟",
}

func (b *Bot) handleWeatherCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

}

func (b *Bot) handleTranslateCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

}

func (b *Bot) handleRemindCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

}

func (b *Bot) handleGameCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

}
