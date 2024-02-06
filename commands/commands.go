package commands

import (
	"RankCheck/embeds"
	"RankCheck/globals"
	"RankCheck/notifs"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mattishere/goverwatch"
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

			authorId := elements[0]

			if i.Member.User.ID != authorId {
				return
			}

			tag := elements[1]

			tagElements := strings.Split(tag, "-")
			name := tagElements[0]
			discriminator, _ := strconv.Atoi(tagElements[1])

			platformString := elements[2]

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})
			if err != nil {
				notifs.Error("Error sending defer response: " + err.Error())
			}

			var stats data.Stats
			if cached, _, exists := globals.StatsCache.Get(tag); exists {
				stats = cached.(data.Stats)
			} else {
				stats, err := goverwatch.GetStats(name, discriminator)
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
				var platform data.Platform
				if platformString == "pc" {
					platform = stats.PC
					embed.Author = &discordgo.MessageEmbedAuthor{
						Name: "PC Stats",
					}
				} else {
					platform = stats.Console
					embed.Author = &discordgo.MessageEmbedAuthor{
						Name: "Console Stats",
					}
					isConsole = true
				}

				embed.Description = stats.Profile.Title

				// FIX, I HAVE NO RANKS AND I STILL GET THE RANKS EXCEPTION!!! BUG IN GOVERWATCH
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
										Value: authorId + ":" + tag + ":pc",
										Emoji: discordgo.ComponentEmoji{
											Name: "⌨️",
										},
										Default: !isConsole,
									},
									{
										Label: "Console",
										Value: authorId + ":" + tag + ":console",
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
	}
}
