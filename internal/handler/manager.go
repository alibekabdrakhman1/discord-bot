package handler

import (
	"cloud.google.com/go/translate"
	"fmt"
	"github.com/alibekabdrakhman1/discord-bot/internal/transport"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"strings"
)

type Manager struct {
	Game      IBotHandler
	Help      IBotHandler
	Poll      IBotHandler
	Weather   IBotHandler
	Translate IBotHandler
	logger    *zap.SugaredLogger
}

func NewManager(translateClient *translate.Client, weatherClient *transport.WeatherClient) *Manager {
	return &Manager{
		Game:      NewGameHandler(),
		Help:      NewHelpHandler(),
		Poll:      NewPollHandler(),
		Weather:   NewWeatherHandler(weatherClient),
		Translate: NewTranslateHandler(translateClient),
	}
}

func (h *Manager) MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	args := strings.Split(message.Content, " ")
	fmt.Println(args)

	switch args[0] {
	case "!help":
		go h.Help.Handle(session, message)
	case "!game":
		go h.Game.Handle(session, message)
	case "!poll":
		go h.Poll.Handle(session, message)
	case "!weather":
		h.Weather.Handle(session, message)
	case "!translate":
		go session.ChannelMessageSend(message.ChannelID, "translate is unavailable")
	}

}

type IBotHandler interface {
	Handle(session *discordgo.Session, message *discordgo.MessageCreate)
}
