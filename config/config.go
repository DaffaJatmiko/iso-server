package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig   `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
	Redis    RedisConfig    `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
	SMTP     SMPTConfig     `mapstructure:",squash"`
}

type ServerConfig struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"DATABASE_DSN"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"REDIS_ADDR"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

type SMPTConfig struct {
	Host     string `mapstructure:"SMTP_HOST"`
	Port     int    `mapstructure:"SMTP_PORT"`
	Username string `mapstructure:"SMTP_USERNAME"`
	Password string `mapstructure:"SMTP_PASSWORD"`
	From     string `mapstructure:"SMTP_FROM"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	log.Println("Config file read successfully")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshaling config: %v", err)
		return nil, err
	}

	log.Printf("Loaded configuration: %+v", config)

	return &config, nil
}
