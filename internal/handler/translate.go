package handler

import (
	"cloud.google.com/go/translate"
	"github.com/bwmarrin/discordgo"
)

type TranslateHandler struct {
	client *translate.Client
}

func NewTranslateHandler(client *translate.Client) *TranslateHandler {
	return &TranslateHandler{
		client: client,
	}
}

func (h *TranslateHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	//TODO implement me
	panic("implement me")
}
