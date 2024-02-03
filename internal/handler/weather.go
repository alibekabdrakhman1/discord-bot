package handler

import (
	"github.com/alibekabdrakhman1/discord-bot/internal/transport"
	"github.com/bwmarrin/discordgo"
)

type WeatherHandler struct {
	client *transport.WeatherClient
}

func NewWeatherHandler(client *transport.WeatherClient) *WeatherHandler {
	return &WeatherHandler{
		client: client,
	}
}

func (h *WeatherHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	//TODO implement me
	panic("implement me")
}
