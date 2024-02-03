package handler

import (
	"github.com/bwmarrin/discordgo"
)

type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

func (h *GameHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	//TODO implement me
	panic("implement me")
}
