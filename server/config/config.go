package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config struct stores entire project configurations
type Config struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	APPConfig      APPConfig      `mapstructure:"app"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
}

// ServerConfig has only server specific configuration
type ServerConfig struct {
	ListenAddr   string        `mapstructure:"listenAddr"`
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	Env          string        `mapstructure:"env"`
}

// APPConfig contains api package related configurations
type APPConfig struct {
	DatabaseConfig   DatabaseConfig
	TaskConfig       ServiceConfig `mapstructure:"task"`
	AdditionalConfig Additional    `mapstructure:"additional"`
}

type ServiceConfig struct {
	DBName string `mapstructure:"dbName"`
}

type Additional struct {
}

type PsqlConfig struct {
	DbConfig DatabaseConfig
	DbName   string
}

// DatabaseConfig contains mongodb related configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"sslMode"`
}

// GetConfig returns entire project configuration
func GetConfig() *Config {
	return GetConfigFromFile("default")
}

// GetConfigFromFile returns configuration from specific file object
func GetConfigFromFile(fileName string) *Config {
	if fileName == "" {
		fileName = "default"
	}

	// looking for filename `default` inside `src/server` dir with `.toml` extension
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath("../../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
		os.Exit(1)
	}
	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	config.APPConfig.DatabaseConfig = config.DatabaseConfig
	return config
}
