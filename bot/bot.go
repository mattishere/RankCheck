package bot

import (
	"RankCheck/commands"
	"RankCheck/listeners"
	"RankCheck/notifs"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func StartBot() {
	var token string
	if os.Getenv("PRODUCTION") == "true" {
		token = os.Getenv("TOKEN")
	} else {
		token = os.Getenv("DEV_TOKEN")
		notifs.Warning("Bot is running in development mode")
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		notifs.Error("Error creating Discord session: " + err.Error())
	}

	listeners.RegisterListeners(s)
	commands.RegisterCommands(s)

	if err != nil {
		notifs.Error("Error creating Application Commands: " + err.Error())
	}

	err = s.Open()
	if err != nil {
		notifs.Error("Error opening Discord session: " + err.Error())
	}

	notifs.System("Bot has been started")
	notifs.Normal("TIP: You can press CTRL+C to stop the bot")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sc

	notifs.System("Bot has been stopped")
	s.Close()
}
