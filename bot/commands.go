package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
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
		"!weather [city] - Get current weather information\n" +
		"!translate [language code] [text] - Translate text to the specified language\n" +
		"!remind [time] [message] - Set a reminder\n" +
		"!game - Play a text-based game\n"
	b.Session.ChannelMessageSend(channelID, helpMessage)
}
func (b *Bot) handlePollCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	words := strings.Fields(message.Content)

	if len(words) < 3 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !poll [question]? -[option1] -[option2] ...")
		return
	}

	optionsString := strings.Join(words[1:], " ")

	questionIndex := strings.Index(optionsString, "?")
	var question string
	var options string
	if questionIndex != -1 {
		question = optionsString[:questionIndex+1]
		options = optionsString[questionIndex+1:]
	}
	fmt.Println("qstr", question, "ostr", options)

	optionsList := strings.Split(options, "-")

	var filteredOptions []string
	for _, option := range optionsList {
		trimmedOption := strings.TrimSpace(option)
		if trimmedOption != "" {
			filteredOptions = append(filteredOptions, trimmedOption)
		}
	}

	if len(filteredOptions) < 2 {
		session.ChannelMessageSend(message.ChannelID, "You need at least two options to create a poll.")
		return
	}

	pollMessage := "Poll: " + question + "\n"
	for i, option := range filteredOptions {
		pollMessage += fmt.Sprintf("%d. %s\n", i+1, option)
	}

	pollMessageSent, err := session.ChannelMessageSend(message.ChannelID, pollMessage)
	if err != nil {
		log.Printf("Error sending poll message: %v", err)
		return
	}

	for i := range filteredOptions {
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
	words := strings.Fields(message.Content)
	if len(words) < 2 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !weather [city]")
		return
	}
	city := words[1]
	weather, err := b.Weather.GetWeatherData(city)
	if err != nil {
		log.Printf("Error getting weather information: %v", err)
		session.ChannelMessageSend(message.ChannelID, "Error getting weather information. Please try again later.")
		return
	}
	response := fmt.Sprintf("Weather information for %s:\n", city)
	response += fmt.Sprintf("Temperature: %.2f°C\n", weather.Main.Temp)
	response += fmt.Sprintf("Humidity: %.2f%%\n", weather.Main.Humidity)
	if len(weather.Weather) > 0 {
		response += fmt.Sprintf("Description: %s\n", weather.Weather[0].Description)
	}
	session.ChannelMessageSend(message.ChannelID, response)
}

func (b *Bot) handleTranslateCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

}
func (b *Bot) handleRemindCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	words := strings.Fields(message.Content)
	if len(words) < 4 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !remind [time] [user] [message]")
		return
	}

	timeStr, userStr, messageText := words[1], words[2], strings.Join(words[3:], " ")

	// Parse time
	duration, err := time.ParseDuration(timeStr)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Invalid time format. Use something like '10s', '5m', '1h', etc.")
		return
	}

	// Log input parameters for debugging
	log.Printf("Time: %s, User: %s, Message: %s", timeStr, userStr, messageText)

	// Get user ID
	user, err := session.User(userStr)
	if err != nil {
		log.Printf("Error getting user info: %v", err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error: %v", err))
		return
	}

	// Calculate trigger time
	triggerAt := time.Now().Add(duration)

	// Add reminder
	b.ReminderManager.AddReminder(user.ID, messageText, triggerAt)

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Reminder set for %s in %s: %s", user.Username, timeStr, messageText))
}

func (b *Bot) handleGameCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

}
