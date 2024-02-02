package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// PollCommand represents a command that creates a poll in a Discord channel.
type PollCommand struct{}

// Execute implements the Command interface for the PollCommand.
// It parses the message content to create a poll with a question and options, and adds reactions for voting.
func (p *PollCommand) Execute(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	// Split the words in the message content
	words := strings.Fields(message.Content)

	// Check if there are enough arguments to create a poll
	if len(words) < 3 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !poll [question]? -[option1] -[option2] ...")
		return
	}
	// Join the words to create a string with the question and options
	optionsString := strings.Join(words[1:], " ")

	// Find the index of the question mark (?) to separate the question and options
	questionIndex := strings.Index(optionsString, "?")
	var question string
	var options string
	if questionIndex != -1 {
		question = optionsString[:questionIndex+1]
		options = optionsString[questionIndex+1:]
	}

	// Split the options string into a list of options
	optionsList := strings.Split(options, "-")

	// Filter out empty options
	var filteredOptions []string
	for _, option := range optionsList {
		trimmedOption := strings.TrimSpace(option)
		if trimmedOption != "" {
			filteredOptions = append(filteredOptions, trimmedOption)
		}
	}

	// Check if there are at least two options to create a poll
	if len(filteredOptions) < 2 {
		session.ChannelMessageSend(message.ChannelID, "You need at least two options to create a poll.")
		return
	}

	// Create the poll message with question and options
	pollMessage := "Poll: " + question + "\n"
	for i, option := range filteredOptions {
		pollMessage += fmt.Sprintf("%d. %s\n", i+1, option)
	}

	// Send the poll message to the Discord channel
	pollMessageSent, err := session.ChannelMessageSend(message.ChannelID, pollMessage)
	if err != nil {
		log.Printf("Error sending poll message: %v", err)
		return
	}
	// Define emoji numbers for reactions
	emojiNumbers := []string{
		"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟",
	}

	// Add reactions for voting to the poll message
	for i := range filteredOptions {
		err := session.MessageReactionAdd(message.ChannelID, pollMessageSent.ID, emojiNumbers[i])
		if err != nil {
			log.Printf("Error adding reaction: %v", err)
			return
		}
	}

}
