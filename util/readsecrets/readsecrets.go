package util

import (
	"encoding/json"
	"os"

	envflags "github.com/harunadess/average-warthunder-player/util/envflags"
	logger "github.com/harunadess/average-warthunder-player/util/logger"
)

var AyameSecrets struct {
	Token              string   `json:"token"`
	DevChannelID       string   `json:"devChannelID"`
	Substrings         []string `json:"substrings"`
	Matchstrings       []string `json:"matchStrings"`
	ChannelIDWhitelist []string `json:"channelIDWhitelist"`
}

func ReadSecrets() {
	fPath := envflags.ConfigPath
	readConfig(fPath)
}

func readConfig(fPath *string) {
	const maxJSONBytes int = 2048

	file, err := os.Open(*fPath)
	if err != nil {
		logger.Fatal("readConfig: ", err)
	}
	defer file.Close()

	data := make([]byte, maxJSONBytes)
	count, err := file.Read(data)
	if err != nil {
		logger.Fatal("readConfig: ", err)
	}

	err = json.Unmarshal(data[:count], &AyameSecrets)
	if err != nil {
		logger.Fatal("readConfig: ", err)
	}
}
