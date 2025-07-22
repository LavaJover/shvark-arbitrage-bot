package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type ArbitrageBotConfig struct {
	Env string 		   `yaml:"env"`
	GRPCServer 		   `yaml:"grpc_server"`
	ArbitrageBotDB 	   `yaml:"arbitrage_db"`
	LogConfig 		   `yaml:"log_config"`
	BotToken	string `yaml:"bot_token"`
	SSOService			`yaml:"sso-service"`
	KafkaService 		`yaml:"kafka-service"`
	AuthzService		`yaml:"authz-service"`
}

type AuthzService struct {
	Host string `yaml:"host"`
	Port string	`yaml:"port"`
}

type KafkaService struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type SSOService struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type GRPCServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ArbitrageBotDB struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	LogLevel 	string 	`yaml:"log_level"`
	LogFormat 	string 	`yaml:"log_format"`
	LogOutput 	string 	`yaml:"log_output"`
}

func MustLoad() *ArbitrageBotConfig {

	// Processing env config variable and file
	configPath := os.Getenv("ARBITRAGE_BOT_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("ARBITRAGE_BOT_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg ArbitrageBotConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}