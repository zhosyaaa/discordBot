package handlers

import "github.com/bwmarrin/discordgo"

// Command is an interface that defines the Execute method for handling Discord commands.
type Command interface {
	// Execute handles the execution of the command.
	// It takes the Discord session, the message that triggered the command, and any command arguments.
	Execute(session *discordgo.Session, message *discordgo.MessageCreate, args []string)
}
