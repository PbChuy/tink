package services

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	*discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate) *discordgo.InteractionResponse
}

var Discord *discordgo.Session
var Commands = make(map[string]*Command)
var Registerdcommands = make([]*discordgo.ApplicationCommand, len(Commands))

func ConnectDiscord(events []interface{}) {
	env := os.Getenv("ENV")
	registerCmds := os.Getenv("REGISTER_CMDS")
	Discord, _ = discordgo.New(os.Getenv("TOKEN"))

	//setting up session
	Discord.Identify.Intents = discordgo.Intent(discordgo.IntentsAllWithoutPrivileged)
	Discord.StateEnabled = true
	Discord.State.MaxMessageCount = 5000

	for _, h := range events {
		Discord.AddHandler(h)
	}

	err := Discord.Open()
	if err != nil {
		panic(err)
	}

	//regesiter commands
	if registerCmds == "true" {
		if env == "prod" {
			println("regestering commands")
			RegisterCmds(Discord, "")
		} else {
			// TODO
		}
	}
}
func RegisterCmds(s *discordgo.Session, g string) {
	//loop
	for _, v := range Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, g, v.ApplicationCommand)
		if err != nil {
			println("failed to register command", v.Name)
			panic(err)
		}
		println("registered command", v.Name)
	}
}
func DisconnectDiscord() {
	Discord.Close()
	println("disconnected from Discord :)")
}
