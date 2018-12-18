package embed

import "github.com/bwmarrin/discordgo"

func getCredits() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Color: 0x005682,
		Type:  "A8out",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Vriska8ot",
				Value:  "Created by \\🐙\\🐙#0413 ( http://oct2pus.tumblr.com/ )\nVriska8ot uses the 'discordgo' library\n( https://github.com/bwmarrin/discordgo/ )",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Special Thanks",
				Value:  "Avatar By mjölk#8323 ( http://cosmic-rumpus.tumblr.com/ )\nEmojis by Dzuk#1671 ( https://noct.zone/ )",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Disclaimer",
				Value:  "Vriska8ot uses **Mutant Standard Emoji** (https://mutant.tech)\n**Mutant Standard Emoji** are licensed under CC-BY-NC-SA 4.0 (https://creativecommons.org/licenses/by-nc-sa/4.0/)",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://raw.githubusercontent.com/oct2pus/vriskabot/master/art/avatar.png",
		},
	}

	return embed
}
