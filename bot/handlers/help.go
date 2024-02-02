package handlers

import "github.com/bwmarrin/discordgo"

// HelpCommand represents a command that provides information about available commands.
type HelpCommand struct{}

// Execute implements the Command interface for the HelpCommand.
// It sends a help message containing a list of available commands to the Discord channel.
func (h *HelpCommand) Execute(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	// Define the help message with a list of available commands
	helpMessage := "Available commands:\n" +
		"!help - Show this help message\n" +
		"!poll [question] [option1] [option2] ... - Create a poll\n" +
		"!weather [city] - Get current weather information\n" +
		"!translate [language code] [text] - Translate text to the specified language\n" +
		"!language - Get languages to translate\n"

	// Send the help message to the Discord channel
	session.ChannelMessageSend(message.ChannelID, helpMessage)
}
