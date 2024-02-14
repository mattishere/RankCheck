package commands

import (
	"RankCheck/embeds"
	"RankCheck/globals"
	"RankCheck/notifs"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mattishere/goverwatch"
	"github.com/mattishere/goverwatch/data"
)

var (
	minVal       = 3
	maxVal       = 12
	StatsCommand = Command{
		Name: "stats",
		Command: discordgo.ApplicationCommand{
			Name:        "stats",
			Description: "Get an Overwatch 2 player's stats",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "The name of the player (e.g. MattHere)",
					MinLength:   &minVal,
					MaxLength:   maxVal,
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "tag",
					Description: "The tag/discriminator of the player (e.g. 2211)",
					Required:    true,
				},
			},
		},
		Execute: statsCommand,
	}
)

func statsCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, cooldown, isOnCooldown := globals.Cooldowns.Get(i.Member.User.ID)
	if isOnCooldown {

		embed := embeds.EmbedWithPFP(s)
		embed.Title = "Slow down!"
		embed.Description = fmt.Sprintf("You're still in queue! You're getting into a match <t:%d:R>", cooldown.Unix())

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	opts := i.ApplicationCommandData().Options
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
	for _, opt := range opts {
		options[opt.Name] = opt
	}

	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "*Diving the stats...*",
			},
		},
	)

	if err != nil {
		notifs.Error("Error responding to interaction: " + err.Error())
	}

	globals.Cooldowns.Set(i.Member.User.ID, nil)

	name := options["name"].StringValue()
	discriminator := int(options["tag"].IntValue())

	tag := fmt.Sprintf("%s-%d", name, discriminator)

	var stats data.Stats
	if cached, _, exists := globals.StatsCache.Get(tag); exists {
		stats = cached.(data.Stats)
	} else {
		stats, err = goverwatch.GetStats(name, discriminator)
		if err != nil {
			embed := embeds.ErrorEmbedWithPFP(s)
			embed.Title = "Error"
			embed.Description = "An error occurred while fetching the stats."
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &globals.Empty,
				Embeds: &[]*discordgo.MessageEmbed{
					embed,
				},
			})
			notifs.Error(err.Error())
			return
		}
		if !stats.Profile.Exists {
			embed := embeds.EmbedWithPFP(s)
			embed.Title = "Player not found"
			embed.Description = "The player you are looking for does not exist.\nConsider checking the spelling of the name and the tag."

			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &globals.Empty,
				Embeds: &[]*discordgo.MessageEmbed{
					embed,
				},
			})

			if err != nil {
				notifs.Error("Error editing response: " + err.Error())
			}

			return
		} else {
			globals.StatsCache.Set(tag, stats)
		}
	}

	embed := embeds.Embed()
	embed.Title = stats.Profile.Name + "'s stats"
	embed.URL = stats.Profile.URL
	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: stats.Profile.ProfilePicture}

	var isConsole bool

	if stats.Profile.IsPrivate {
		embed.Description = stats.Profile.Title + "\n\n*This player's profile is private.*"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &globals.Empty,
			Embeds: &[]*discordgo.MessageEmbed{
				embed,
			},
		})

		return
	} else {
		platform := stats.PC
		embed.Author = &discordgo.MessageEmbedAuthor{
			Name: "PC Stats",
		}

		if !platform.HasRanks {
			platform = stats.Console
			embed.Author = &discordgo.MessageEmbedAuthor{
				Name: "Console Stats",
			}
			isConsole = true
		}

		embed.Description = stats.Profile.Title

		if platform.HasRanks {
			embed.Fields = getRanks(platform)
		} else {
			embed.Description += "\n\nThis player has no ranks on this platform."
		}
	}

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &globals.Empty,
		Embeds: &[]*discordgo.MessageEmbed{
			embed,
		},
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID:    "platform",
						Placeholder: "Select platform",
						Options: []discordgo.SelectMenuOption{
							{
								Label: "PC",
								Value: i.Member.User.ID + ":" + tag + ":pc",
								Emoji: discordgo.ComponentEmoji{
									Name: "⌨️",
								},
								Default: !isConsole,
							},
							{
								Label: "Console",
								Value: i.Member.User.ID + ":" + tag + ":console",
								Emoji: discordgo.ComponentEmoji{
									Name: "🎮",
								},
								Default: isConsole,
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		notifs.Error("Error editing response: " + err.Error())
	}
}

func roleIsRanked(rank data.Rank) bool {
	return rank.Rank != ""
}

func getRanks(platform data.Platform) []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}

	if roleIsRanked(platform.Ranks.Tank) {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Tank",
			Value:  fmt.Sprintf("%s %d", platform.Ranks.Tank.Rank, platform.Ranks.Tank.Division),
			Inline: true,
		})
	}
	if roleIsRanked(platform.Ranks.DPS) {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "DPS",
			Value:  fmt.Sprintf("%s %d", platform.Ranks.DPS.Rank, platform.Ranks.DPS.Division),
			Inline: true,
		})
	}
	if roleIsRanked(platform.Ranks.Support) {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Support",
			Value:  fmt.Sprintf("%s %d", platform.Ranks.Support.Rank, platform.Ranks.Support.Division),
			Inline: true,
		})
	}
	if roleIsRanked(platform.Ranks.OpenQueue) {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Open Queue",
			Value:  fmt.Sprintf("%s %d", platform.Ranks.OpenQueue.Rank, platform.Ranks.OpenQueue.Division),
			Inline: true,
		})
	}

	return fields
}
