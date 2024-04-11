package globals

var (
	Tanks = map[string]Hero{
		"dva": {
			Name:  "D.Va",
			Emoji: "🐰",
		},
		"doomfist": {
			Name:  "Doomfist",
			Emoji: "🤜",
		},
		"junker-queen": {
			Name:  "Junker Queen",
			Emoji: "🪓",
		},
		"mauga": {
			Name:  "Mauga",
			Emoji: "🌴",
		},
		"orisa": {
			Name:  "Orisa",
			Emoji: "🐴",
		},
		"ramattra": {
			Name:  "Ramattra",
			Emoji: "🌘",
		},
		"reinhardt": {
			Name:  "Reinhardt",
			Emoji: "🛡️",
		},
		"roadhog": {
			Name:  "Roadhog",
			Emoji: "🐷",
		},
		"sigma": {
			Name:  "Sigma",
			Emoji: "🧠",
		},
		"winston": {
			Name:  "Winston",
			Emoji: "🐵",
		},
		"wrecking-ball": {
			Name:  "Wrecking Ball",
			Emoji: "🐹",
		},
		"zarya": {
			Name:  "Zarya",
			Emoji: "💪",
		},
	}

	DPS = map[string]Hero{
		"ashe": {
			Name:  "Ashe",
			Emoji: "🧨",
		},
		"bastion": {
			Name:  "Bastion",
			Emoji: "🐦",
		},
		"cassidy": {
			Name:  "Cassidy",
			Emoji: "🤠",
		},
		"echo": {
			Name:  "Echo",
			Emoji: "🌌",
		},
		"genji": {
			Name:  "Genji",
			Emoji: "⚔️",
		},
		"hanzo": {
			Name:  "Hanzo",
			Emoji: "🏹",
		},
		"junkrat": {
			Name:  "Junkrat",
			Emoji: "💣",
		},
		"mei": {
			Name:  "Mei",
			Emoji: "❄️",
		},
		"pharah": {
			Name:  "Pharah",
			Emoji: "🚀",
		},
		"reaper": {
			Name:  "Reaper",
			Emoji: "👻",
		},
		"sojourn": {
			Name:  "Sojourn",
			Emoji: "🍁",
		},
		"soldier-76": {
			Name:  "Soldier: 76",
			Emoji: "🔫",
		},
		"sombra": {
			Name:  "Sombra",
			Emoji: "⌨️",
		},
		"symmetra": {
			Name:  "Symmetra",
			Emoji: "🧱",
		},
		"torbjorn": {
			Name:  "Torbjörn",
			Emoji: "🔨",
		},
		"tracer": {
			Name:  "Tracer",
			Emoji: "⚡",
		},
		"widowmaker": {
			Name:  "Widowmaker",
			Emoji: "🕸️",
		},
	}

	Supports = map[string]Hero{
		"ana": {
			Name:  "Ana",
			Emoji: "💤",
		},
		"baptiste": {
			Name:  "Baptiste",
			Emoji: "💉",
		},
		"brigitte": {
			Name:  "Brigitte",
			Emoji: "🛡️",
		},
		"illari": {
			Name:  "Illari",
			Emoji: "☀️",
		},
		"kiriko": {
			Name:  "Kiriko",
			Emoji: "🦊",
		},
		"lifeweaver": {
			Name:  "Lifeweaver",
			Emoji: "🌸",
		},
		"lucio": {
			Name:  "Lúcio",
			Emoji: "🐸",
		},
		"mercy": {
			Name:  "Mercy",
			Emoji: "🪽",
		},
		"moira": {
			Name:  "Moira",
			Emoji: "🧪",
		},
		"zenyatta": {
			Name:  "Zenyatta",
			Emoji: "🙏",
		},
	}
)

type Hero struct {
	Name  string
	Emoji string
}
