package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME     string
	DB_PASSWORD     string
	DB_PORT         string
	DB_HOST         string
	DB_NAME         string
	DB_NAME_TEST    string
	ENCRYCPTION_KEY string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("app/config/config.json", &configuration)
	return configuration
}