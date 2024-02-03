package handler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type HelpHandler struct {
}

func NewHelpHandler() *HelpHandler {
	return &HelpHandler{}
}

func (h *HelpHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	args := strings.Split(message.Content, " ")
	if len(args) == 1 {
		// Send a general help message
		helpMessage := "Welcome to the bot! Here are available commands:\n\n" +
			"**Game Commands:**\n" +
			"!game rock|scissors|paper - Play Rock-Paper-Scissors with the bot.\n\n" +
			"**Weather Commands:**\n" +
			"!weather <location> - Get current weather information for a location."

		session.ChannelMessageSend(message.ChannelID, helpMessage)
	} else if len(args) == 2 {
		switch args[1] {
		case "game":
			session.ChannelMessageSend(message.ChannelID, "Command: !game rock|scissors|paper\nPlay Rock-Paper-Scissors with the bot.")
		case "weather":
			session.ChannelMessageSend(message.ChannelID, "Command: !weather <location>\nGet current weather information for a location.")
		default:
			session.ChannelMessageSend(message.ChannelID, "Unknown command. Use !help to see available commands.")
		}
	} else {
		session.ChannelMessageSend(message.ChannelID, "Invalid usage. Use !help for general help or !help <command> for command-specific help.")
	}
}
