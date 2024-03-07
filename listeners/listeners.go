package listeners

import (
	"RankCheck/commands"
	"RankCheck/notifs"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateGameStatus(0, "Overwatch 2")
	notifs.System("Bot status set")
}

func RegisterListeners(s *discordgo.Session) {
	s.AddHandler(ready)
	s.AddHandler(commands.InteractionCreateListener)

	notifs.System("Listeners registered")
}