package main

import (
	"os"
	"os/signal"
	"syscall"

	discordactions "github.com/harunadess/average-warthunder-player/discord/discordactions"
	discordinit "github.com/harunadess/average-warthunder-player/discord/discordinit"
	"github.com/harunadess/average-warthunder-player/messaging"
	logger "github.com/harunadess/average-warthunder-player/util/logger"
	readsecrets "github.com/harunadess/average-warthunder-player/util/readsecrets"
)

func main() {
	readsecrets.ReadSecrets()

	ayame := discordinit.SetupBot(readsecrets.AyameSecrets.Token)
	logger.Info("starting!")

	messaging.SendMessage(ayame, readsecrets.AyameSecrets.DevChannelID, "'sup")

	err := discordactions.SetActivity(ayame, "playing Realistic Ground Battles")
	if err != nil {
		logger.Error("SetActvitiy: ", err)
	}

	// wait here for ctrl+c or other signal end term
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// cleanly close the discord session
	ayame.Close()
}
