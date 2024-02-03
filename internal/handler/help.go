package handler

import (
	"github.com/bwmarrin/discordgo"
)

type HelpHandler struct {
}

func NewHelpHandler() *HelpHandler {
	return &HelpHandler{}
}

func (h *HelpHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	//TODO implement me
	panic("implement me")
}
