package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Uwusie/dinozalr2-bot/internal/config"
	"github.com/Uwusie/dinozalr2-bot/internal/handlers"
	"github.com/Uwusie/dinozalr2-bot/internal/rabbitmq"
	"github.com/bwmarrin/discordgo"
)

func main() {
	botConfig := config.Load()
	discord, err := discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	commandConfig := config.LoadCommandConfig()
	registerHandlers(discord, commandConfig)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go rabbitmq.ListenForMessages()
	defer rabbitmq.CloseConnection()
	defer rabbitmq.CloseChannel()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func registerHandlers(s *discordgo.Session, c *config.CommandConfig) {
	s.AddHandler(handlers.NewMessageCreateEventHandler(c))
}
