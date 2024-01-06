package notifs

import (
	"fmt"
	"log"
	"os"
)

func Error(msg string) {
	LogMessage("\x1b[31m", "❌", msg)
}

func Warning(msg string) {
	LogMessage("\x1b[33m", "🚧", msg)
}

func System(msg string) {
	LogMessage("\x1b[34m", "📂", msg)
}

func Background(msg string) {
	LogMessage("\x1b[90m", "🧹", msg)
}

func Normal(msg string) {
	LogMessage("\x1b[37m", "📜", msg)
}

func LogMessage(color string, prefix string, message string) {
	if os.Getenv("LOGS") == "true" {
		final := fmt.Sprintf("%s%s | %s%s", color, prefix, message, "\x1b[0m")
		log.Println(final)
	}
}
