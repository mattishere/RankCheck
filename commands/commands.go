package commands

import (
	"RankCheck/notifs"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []Command{
		StatsCommand,
		DiscordCommand,
		LinksCommand,
		InviteCommand,
	}
)

type Command struct {
	Name    string
	Command discordgo.ApplicationCommand
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func RegisterCommands(s *discordgo.Session) {
	var appId string
	if os.Getenv("PRODUCTION") == "true" {
		appId = os.Getenv("APP_ID")
	} else {
		appId = os.Getenv("DEV_APP_ID")
	}

	for _, command := range commands {
		cmd, err := s.ApplicationCommandCreate(appId, "", &command.Command)
		if err != nil {
			panic(err)
		}
		notifs.Background("Registered command /" + cmd.Name)
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
