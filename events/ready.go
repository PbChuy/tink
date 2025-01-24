package events

import "github.com/bwmarrin/discordgo"

func init() {
	Events = append(Events, onReady)
}

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateCustomStatus("fart")
	println("Signed in as " + s.State.User.String())
}
