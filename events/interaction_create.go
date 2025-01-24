package events

import (
	"github.com/PbChuy/tink/commands"
	"github.com/bwmarrin/discordgo"
)

func init() {
	Events = append(Events, onInteractionCreate)

}

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		commands.OnInteraction(s, i)
	default:
		println("unhandled interaction type", string(i.Type))
	}
}
