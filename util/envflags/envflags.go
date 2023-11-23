package envflags

import "flag"

// ConfigPath contains the file path to the main config for the bot
var ConfigPath *string

func init() {
	ConfigPath = flag.String("config", "", "path to config .json file")
	flag.Parse()
}

type Secrets struct {
	Token        string   `json:"token"`
	DevChannelID string   `json:"devChannelID"`
	Substrings   []string `json:"substrings"`
}
