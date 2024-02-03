package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"time"
)

type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

var choices = []string{"rock", "scissors", "paper"}

func (h *GameHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	playerChoice := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(message.Content, "!game ")))

	if !isValidChoice(playerChoice) {
		session.ChannelMessageSend(message.ChannelID, "Invalid choice. Use: `!game rock|scissors|paper`")
		return
	}

	botChoice := getBotChoice()
	result := determineWinner(playerChoice, botChoice)

	response := fmt.Sprintf("You chose %s, I chose %s. %s", playerChoice, botChoice, result)
	session.ChannelMessageSend(message.ChannelID, response)
}

func getBotChoice() string {
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(len(choices))]
}

func isValidChoice(choice string) bool {
	for _, validChoice := range choices {
		if choice == validChoice {
			return true
		}
	}
	return false
}

func determineWinner(playerChoice, botChoice string) string {
	switch {
	case playerChoice == botChoice:
		return "It's a tie!"
	case (playerChoice == "rock" && botChoice == "scissors") ||
		(playerChoice == "scissors" && botChoice == "paper") ||
		(playerChoice == "paper" && botChoice == "rock"):
		return "You win!"
	default:
		return "You lose!"
	}
}
