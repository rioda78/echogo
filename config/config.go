package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	Rdms       string
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("config/config.json", &conf)
	if err != nil {
		return Configuration{}
	}
	return conf
}
