package handlers

import "github.com/bwmarrin/discordgo"

// LanguageCommand represents a command that provides a list of supported languages for translation.
type LanguageCommand struct{}

// Execute implements the Command interface for the LanguageCommand.
// It sends a message containing a list of supported languages and their language codes to the Discord channel.
func (h *LanguageCommand) Execute(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	// Define the message with a list of supported languages and their codes
	languageMessage := "English: en\n" +
		"Spanish: es\n" +
		"French: fr\n" +
		"German: de\n" +
		"Chinese (Simplified): zh-CN" +
		"\nChinese (Traditional): zh-TW" +
		"\nJapanese: ja" +
		"\nKorean: ko" +
		"\nRussian: ru" +
		"\nArabic: ar" +
		"\nHindi: hi" +
		"\nPortuguese: pt" +
		"\nItalian: it" +
		"\nDutch: nl" +
		"\nTurkish: tr"
	// Send the language message to the Discord channel
	session.ChannelMessageSend(message.ChannelID, languageMessage)
}
