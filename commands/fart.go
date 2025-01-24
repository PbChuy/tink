package commands

import (
	"github.com/PbChuy/tink/services"
	"github.com/bwmarrin/discordgo"
)

func init() {
	services.Commands[fartCmd.Name] = &services.Command{
		ApplicationCommand: fartCmd,
		Handler:            handleFart,
	}
}

var fartCmd = &discordgo.ApplicationCommand{
	Type:        discordgo.ChatApplicationCommand,
	Name:        "fart",
	Description: "Fart loud",
}

func handleFart(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	return ContentResponse("fawt for me bbg", false)
}
