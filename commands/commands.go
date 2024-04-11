package commands

import (
	"RankCheck/globals"
	"RankCheck/notifs"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mattishere/goverwatch/data"
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
	if i.Type == discordgo.InteractionApplicationCommand {
		cmdData := i.ApplicationCommandData()

		for _, command := range commands {
			if cmdData.Name == command.Name {
				command.Execute(s, i)
			}
		}
	} else if i.Type == discordgo.InteractionMessageComponent {
		if i.MessageComponentData().CustomID == "platform" {
			elements := strings.Split(i.MessageComponentData().Values[0], ":")

			author := elements[0]

			if i.Member.User.ID != author {
				return
			}

			tag := elements[1]

			var isConsole bool
			if elements[2] == "console" {
				isConsole = true
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})
			if err != nil {
				notifs.Error("Error sending defer response: " + err.Error())
			}

			if cached, _, exists := globals.StatsCache.Get(tag); exists {
				SendStatsResponse(author, tag, cached.(data.Stats), isConsole, s, i)
			}
		}
	}
}
