package handlers

import (
	"discordBot/clients"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// TranslateCommand represents a command for translating text to a specified language.
type TranslateCommand struct {
	Translate *clients.TranslateClient
}

// Execute implements the Command interface for the TranslateCommand.
// It translates the provided text to the specified language and sends the translated text to the Discord channel.
func (t *TranslateCommand) Execute(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	// Split the words in the message content
	words := strings.Fields(message.Content)

	// Check if there are enough arguments to perform translation
	if len(words) < 3 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !translate [language code] [text]")
		return
	}
	// Extract target language and text to translate
	targetLanguage := words[1]
	textToTranslate := strings.Join(words[2:], " ")

	// Perform translation
	translatedText, err := t.Translate.TranslateText(textToTranslate, targetLanguage)
	if err != nil {
		log.Printf("Error translating text: %v", err)
		session.ChannelMessageSend(message.ChannelID, "Error translating text. Please try again later.")
		return
	}

	// Send the translated text to the Discord channel
	response := fmt.Sprintf("Translation to %s:\n%s", targetLanguage, translatedText)
	session.ChannelMessageSend(message.ChannelID, response)
}
