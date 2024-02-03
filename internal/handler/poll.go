package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type PollHandler struct {
}

func NewPollHandler() *PollHandler {
	return &PollHandler{}
}

func (h *PollHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	parts := strings.Split(strings.TrimSpace(strings.TrimPrefix(message.Content, "!poll")), "|")

	if len(parts) < 2 {
		session.ChannelMessageSend(message.ChannelID, "Invalid poll format. Use: `!poll <question> | <option1> | <option2> | ...`")
		return
	}

	question := strings.TrimSpace(parts[0])
	options := make([]string, len(parts)-1)

	for i := 1; i < len(parts); i++ {
		options[i-1] = strings.TrimSpace(parts[i])
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("📊**%s**\n\n", question))
	pollContent := ""
	for i, option := range options {
		emoji := getEmojiForIndex(i)
		pollContent += fmt.Sprintf("%s : %s\n", emoji, option)
	}

	embed := &discordgo.MessageEmbed{
		Description: pollContent,
	}

	m, _ := session.ChannelMessageSendEmbed(message.ChannelID, embed)

	for i := range options {
		emoji := getEmojiForIndex(i)
		err := session.MessageReactionAdd(message.ChannelID, m.ID, emoji)
		if err != nil {
			fmt.Println("Error adding reaction:", err)
		}
	}
}

func getEmojiForIndex(index int) string {
	if index >= 0 && index < 26 {
		return string('🇦' + index)
	}
	return ""
}
