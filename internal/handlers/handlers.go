package handlers

import (
	"strings"

	"github.com/Uwusie/dinozalr2-bot/internal/reactions"
	"github.com/bwmarrin/discordgo"
)

func MessageCreateEventHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "meow") {
		reactions.MeowReaction{}.React(s, m)
	}
}
