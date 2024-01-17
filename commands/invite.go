package commands

import (
	"RankCheck/embeds"
	"RankCheck/globals"

	"github.com/bwmarrin/discordgo"
)

var (
	InviteCommand = Command{
		Name: "invite",
		Command: discordgo.ApplicationCommand{
			Name:        "invite",
			Description: "Invite RankCheck to your server",
		},
		Execute: inviteCommand,
	}
)

func inviteCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := embeds.EmbedWithPFP(s)
	embed.Title = "Invite RankCheck to your server"
	embed.URL = globals.BotInvite
	embed.Description = "Invite RankCheck to your server and get great Overwatch 2 statistics in a matter of seconds!"

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "✉️",
							},
							Label: "Invite RankCheck",
							Style: discordgo.LinkButton,
							URL:   globals.BotInvite,
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
