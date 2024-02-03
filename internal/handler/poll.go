package handler

import (
	"github.com/bwmarrin/discordgo"
)

type PollHandler struct {
}

func NewPollHandler() *PollHandler {
	return &PollHandler{}
}

func (h *PollHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	//TODO implement me
	panic("implement me")
}
