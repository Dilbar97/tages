package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

const confPath = "config.yaml"

type Config struct {
	DB           string `yaml:"db"`
	GrpcPort     int    `yaml:"grpc_port"`
	FileLocation string `yaml:"file_location"`
}

func GetConfigs() Config {
	conf := Config{}

	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatal().Msgf("config file '%s' not exist", confPath)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal().Msg("config file is not correct format")
	}

	return conf
}
