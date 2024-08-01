// local testing
//package config
//
//import (
//	"errors"
//	"github.com/spf13/viper"
//	"log"
//	"strconv"
//)
//
//type Config struct {
//	Server   ServerConfig   `mapstructure:",squash"`
//	Database DatabaseConfig `mapstructure:",squash"`
//	JWT      JWTConfig      `mapstructure:",squash"`
//	SMTP     SMTPConfig     `mapstructure:",squash"`
//}
//
//type ServerConfig struct {
//	Port string `mapstructure:"SERVER_PORT"`
//}
//
//type DatabaseConfig struct {
//	DSN string `mapstructure:"DATABASE_DSN"`
//}
//
//type JWTConfig struct {
//	SecretKey string `mapstructure:"JWT_SECRET_KEY"`
//}
//
//type SMTPConfig struct {
//	Host     string `mapstructure:"SMTP_HOST"`
//	Port     int    `mapstructure:"SMTP_PORT"`
//	Username string `mapstructure:"SMTP_USERNAME"`
//	Password string `mapstructure:"SMTP_PASSWORD"`
//	From     string `mapstructure:"SMTP_FROM"`
//}
//
//func LoadConfig() (*Config, error) {
//	// Set the file name of the configurations file
//	viper.SetConfigFile(".env")
//
//	// Read in environment variables that match
//	viper.AutomaticEnv()
//
//	// Find and read the config file
//	if err := viper.ReadInConfig(); err != nil {
//		log.Printf("Error reading config file: %v", err)
//		return nil, err
//	}
//
//	// Get SMTP_PORT value from environment variables
//	smtpPortStr := viper.GetString("SMTP_PORT")
//	if smtpPortStr == "" {
//		log.Println("Error: SMTP_PORT is not set")
//		return nil, errors.New("SMTP_PORT is not set")
//	}
//
//	// Convert SMTP_PORT to integer
//	port, err := strconv.Atoi(smtpPortStr)
//	if err != nil {
//		log.Printf("Error parsing SMTP_PORT: %v", err)
//		return nil, err
//	}
//
//	config := &Config{
//		Server: ServerConfig{
//			Port: viper.GetString("SERVER_PORT"),
//		},
//		Database: DatabaseConfig{
//			DSN: viper.GetString("DATABASE_DSN"),
//		},
//		JWT: JWTConfig{
//			SecretKey: viper.GetString("JWT_SECRET"),
//		},
//		SMTP: SMTPConfig{
//			Host:     viper.GetString("SMTP_HOST"),
//			Port:     port,
//			Username: viper.GetString("SMTP_USERNAME"),
//			Password: viper.GetString("SMTP_PASSWORD"),
//			From:     viper.GetString("SMTP_FROM"),
//		},
//	}
//
//	log.Printf("Loaded configuration: %+v", config)
//
//	return config, nil
//}

// -------------------------------------------------------------------------------------------------------------------

// // Production Config
package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig   `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
	SMTP     SMTPConfig     `mapstructure:",squash"`
}

type ServerConfig struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"DATABASE_DSN"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"SMTP_HOST"`
	Port     int    `mapstructure:"SMTP_PORT"`
	Username string `mapstructure:"SMTP_USERNAME"`
	Password string `mapstructure:"SMTP_PASSWORD"`
	From     string `mapstructure:"SMTP_FROM"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Printf("Error parsing SMTP_PORT: %v", err)
		return nil, err
	}

	config := &Config{
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: DatabaseConfig{
			DSN: os.Getenv("DATABASE_DSN"),
		},
		JWT: JWTConfig{
			SecretKey: os.Getenv("JWT_SECRET"),
		},
		SMTP: SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     port,
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			From:     os.Getenv("SMTP_FROM"),
		},
	}

	log.Printf("Loaded configuration: %+v", config)

	return config, nil
}

// ----------------------------------------------------------------------
//package config
//
//import (
//	"github.com/spf13/viper"
//	"log"
//)
//
//type Config struct {
//	Server   ServerConfig   `mapstructure:",squash"`
//	Database DatabaseConfig `mapstructure:",squash"`
//	Redis    RedisConfig    `mapstructure:",squash"`
//	JWT      JWTConfig      `mapstructure:",squash"`
//	SMTP     SMPTConfig     `mapstructure:",squash"`
//}
//
//type ServerConfig struct {
//	Port string `mapstructure:"SERVER_PORT"`
//}
//
//type DatabaseConfig struct {
//	DSN string `mapstructure:"DATABASE_DSN"`
//}
//
//type RedisConfig struct {
//	Addr     string `mapstructure:"REDIS_ADDR"`
//	Password string `mapstructure:"REDIS_PASSWORD"`
//	DB       int    `mapstructure:"REDIS_DB"`
//}
//
//type JWTConfig struct {
//	SecretKey string `mapstructure:"JWT_SECRET_KEY"`
//}
//
//type SMPTConfig struct {
//	Host     string `mapstructure:"SMTP_HOST"`
//	Port     int    `mapstructure:"SMTP_PORT"`
//	Username string `mapstructure:"SMTP_USERNAME"`
//	Password string `mapstructure:"SMTP_PASSWORD"`
//	From     string `mapstructure:"SMTP_FROM"`
//}
//
//func LoadConfig() (*Config, error) {
//	viper.SetConfigFile(".env")
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err != nil {
//		log.Printf("Error reading config file: %v", err)
//		return nil, err
//	}
//
//	log.Println("Config file read successfully")
//
//	var config Config
//	if err := viper.Unmarshal(&config); err != nil {
//		log.Printf("Error unmarshaling config: %v", err)
//		return nil, err
//	}
//
//	log.Printf("Loaded configuration: %+v", config)
//
//	return &config, nil
//}
