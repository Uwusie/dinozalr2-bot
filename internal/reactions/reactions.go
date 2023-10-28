package reactions

import (
	"fmt"
	"strings"

	"github.com/Uwusie/dinozalr2-bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

type MessageCreateReaction interface {
	React(s *discordgo.Session, m *discordgo.MessageCreate)
}

type MeowReaction struct{}

func (MeowReaction) React(s *discordgo.Session, m *discordgo.MessageCreate, c *config.CommandConfig) {
	message := strings.Repeat("Meow! ", c.MeowCount)
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		fmt.Println("error sending message,", err)
	}
}
