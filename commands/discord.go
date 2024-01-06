package commands

import (
	"RankCheck/embeds"
	"RankCheck/globals"

	"github.com/bwmarrin/discordgo"
)

var (
	DiscordCommand = Command{
		Name: "discord",
		Command: discordgo.ApplicationCommand{
			Name:        "discord",
			Description: "Get an invite to our Discord server",
		},
		Execute: discordCommand,
	}
)

func discordCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := embeds.EmbedWithPFP(s)
	embed.Title = "Join our Discord!"
	embed.URL = globals.DiscordInvite
	embed.Description = "Join our Discord server to report bugs, get help and join our community!"

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
