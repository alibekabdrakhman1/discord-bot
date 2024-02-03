package applicator

import (
	"cloud.google.com/go/translate"
	"context"
	"github.com/alibekabdrakhman1/discord-bot/internal/config"
	"github.com/alibekabdrakhman1/discord-bot/internal/handler"
	"github.com/alibekabdrakhman1/discord-bot/internal/transport"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func New(logger *zap.SugaredLogger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	session, err := discordgo.New("Bot " + a.config.Discord.Token)
	if err != nil {
		a.logger.Fatal(err)
	}
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	translateClient, err := translate.NewClient(ctx, option.WithAPIKey(a.config.Translate.APIKey))
	weatherClient := transport.NewWeatherClient(a.config)

	commandHandler := handler.NewManager(translateClient, weatherClient)

	session.AddHandler(commandHandler.MessageCreate)

	err = session.Open()
	if err != nil {
		a.logger.Fatal(err)
	}
	defer session.Close()

	a.logger.Info("Bot is online")
	gracefullyShutdown(cancel)

}

func gracefullyShutdown(c context.CancelFunc) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-shutdownSignal
}
