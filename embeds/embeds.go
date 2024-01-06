package embeds

import "github.com/bwmarrin/discordgo"

var (
	Empty      = ""
	Footer     = "Made by @matthereatm | Join our /discord for support"
	Color      = 0xDCEBEA
	ErrorColor = 0xEB5F5F
)

func Embed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
		Color: Color,
	}
}

func EmbedWithPFP(s *discordgo.Session) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: s.State.User.AvatarURL("256"),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
		Color: Color,
	}
}

func ErrorEmbed() *discordgo.MessageEmbed {
	embed := Embed()
	embed.Color = ErrorColor
	return embed
}

func ErrorEmbedWithPFP(s *discordgo.Session) *discordgo.MessageEmbed {
	embed := EmbedWithPFP(s)
	embed.Color = ErrorColor
	return embed
}
