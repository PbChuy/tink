package commands

import (
	"time"

	"github.com/PbChuy/tink/lib"
	"github.com/PbChuy/tink/services"
	"github.com/bwmarrin/discordgo"
)

func OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// start a timer to track how long the command takes
	start := time.Now()

	data := i.ApplicationCommandData()
	cmd, ok := services.Commands[data.Name]
	if !ok {
		s.InteractionRespond(i.Interaction, ContentResponse("<:error:1228053905590718596> **Error:** Command does not exist", true))
		return
	}

	resp := cmd.Handler(s, i)
	if resp != nil {
		s.InteractionRespond(i.Interaction, resp)
	} else {
		println("Something went wrong while processing a command: %s", i.ApplicationCommandData().Name)
		errMessage := "<:error:1228053905590718596> **Error:** Something went wrong while processing the command"
		s.InteractionRespond(i.Interaction, ContentResponse(errMessage, true))
		return
	}

	// stop the timer
	end := time.Now()
	lib.CmdRun(s, i, end.Sub(start))
}

func LoadingResponse() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}
}

func ContentResponse(c string, e bool) *discordgo.InteractionResponse {
	d := &discordgo.InteractionResponseData{
		Content:         c,
		AllowedMentions: new(discordgo.MessageAllowedMentions),
	}
	if e {
		d.Flags = discordgo.MessageFlagsEphemeral
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: d,
	}
}

func EmbedResponse(e *discordgo.MessageEmbed, f bool) *discordgo.InteractionResponse {
	d := &discordgo.InteractionResponseData{
		Embeds:          []*discordgo.MessageEmbed{e},
		AllowedMentions: new(discordgo.MessageAllowedMentions),
	}
	if f {
		d.Flags = discordgo.MessageFlagsEphemeral
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: d,
	}
}
