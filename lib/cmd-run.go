package lib

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CmdRun(s *discordgo.Session, i *discordgo.InteractionCreate, d time.Duration) {
	data := i.ApplicationCommandData()
	fmt.Printf("Command executed:\n\tcmd: %s, guild: %s, user: %s, took: %s", data.Name, i.GuildID, i.Member.User.ID, d.String())
}
