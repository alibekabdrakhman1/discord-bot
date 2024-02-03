package handler

import (
	"fmt"
	"github.com/alibekabdrakhman1/discord-bot/internal/transport"
	"github.com/bwmarrin/discordgo"
	"strings"
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
	args := strings.Split(message.Content, " ")

	if len(args) != 2 {
		session.ChannelMessageSend(message.ChannelID, "Invalid usage. Use !help for general help or !help <command> for command-specific help.")
	}

	weatherData, err := h.client.GetWeather(args[1])
	fmt.Println()
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, err.Error())
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Current Weather Information",
		Description: fmt.Sprintf("Temperature: %.2f°C\nFeels like: %.2f°C\nCondition: %s", weatherData.Current.TempC, weatherData.Current.FeelsLikeC, weatherData.Current.Condition.Text),
		Color:       0x3498db,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Location",
				Value:  fmt.Sprintf("%s, %s, %s", weatherData.Location.Name, weatherData.Location.Region, weatherData.Location.Country),
				Inline: true,
			},
			{
				Name:   "Wind Speed",
				Value:  fmt.Sprintf("%.2f KPH", weatherData.Current.WindKPH),
				Inline: true,
			},
			{
				Name:   "Humidity",
				Value:  fmt.Sprintf("%d%%", weatherData.Current.Humidity),
				Inline: true,
			},
		},
	}
	session.ChannelMessageSendEmbed(message.ChannelID, embed)
}
