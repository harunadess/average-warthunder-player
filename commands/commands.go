package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/harunadess/average-warthunder-player/messaging"
	logger "github.com/harunadess/average-warthunder-player/util/logger"
	readsecrets "github.com/harunadess/average-warthunder-player/util/readsecrets"
)

// OnMessageCreate is a handler for the discord event MessageCreate
func OnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	for _, whiteListed := range readsecrets.AyameSecrets.ChannelIDWhitelist {
		if strings.Contains(message.ChannelID, whiteListed) {
			return
		}
	}

	for _, substr := range readsecrets.AyameSecrets.Substrings {
		if strings.Contains(strings.TrimSpace(strings.ToLower(message.Content)), substr) {
			sendMessage(session, message)
			return
		}
	}

	for _, token := range strings.Split(message.Content, " ") {
		for _, matcher := range readsecrets.AyameSecrets.Matchstrings {
			if strings.TrimSpace(strings.ToLower(token)) == matcher {
				sendMessage(session, message)
				return
			}
		}
	}
}

func sendMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	logger.Info("sending response to: ", message.ID)
	messaging.Reply(session, message.ChannelID, message.Reference(), "https://harunadess.com/silence.png")
}
