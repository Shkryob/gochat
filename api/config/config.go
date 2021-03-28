package config

import (
	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/json"
	"github.com/labstack/gommon/log"
)

type Configuration struct {
	AuthRedirectUrl string
	MaxChatParticipants int
}

func ReadConfig() Configuration {
	configuration := Configuration{}

	//Base configuration values common for all envs
	goconfig.File = "./config/config.base.json"

	err := goconfig.Parse(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
