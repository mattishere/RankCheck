package commands

import (
	"RankCheck/globals"
	"RankCheck/notifs"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []Command{
		StatsCommand,
		DiscordCommand,
	}
)

type Command struct {
	Name    string
	Command discordgo.ApplicationCommand
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func RegisterCommands(s *discordgo.Session) {
	for _, command := range commands {
		_, err := s.ApplicationCommandCreate(globals.AppID, "", &command.Command)
		if err != nil {
			panic(err)
		}
	}
	notifs.System("Commands registered")
}

func InteractionCreateListener(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmdData := i.ApplicationCommandData()
	for _, command := range commands {
		if cmdData.Name == command.Name {
			command.Execute(s, i)
		}
	}
}
