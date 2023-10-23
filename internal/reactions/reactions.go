package reactions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MessageCreateReaction interface {
	React(s *discordgo.Session, m *discordgo.MessageCreate)
}

type MeowReaction struct{}

func (MeowReaction) React(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Meow!")
	if err != nil {
		fmt.Println("error sending message,", err)
	}
}
