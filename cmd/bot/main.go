package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Uwusie/dinozalr2-bot/internal/config"
	"github.com/Uwusie/dinozalr2-bot/internal/handlers"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config := config.Load()
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	registerHandlers(discord)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func registerHandlers(s *discordgo.Session) {
	s.AddHandler(handlers.MessageCreateEventHandler)
}
