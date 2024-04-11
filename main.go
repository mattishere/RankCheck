package main

import (
	"RankCheck/bot"
	"RankCheck/globals"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	go globals.UpdateCaches()

	bot.StartBot()
}
