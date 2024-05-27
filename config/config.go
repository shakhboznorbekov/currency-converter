package config

import (
	"log"
	"os"
)

type Config struct {
	Server      ServerConfig      `yaml:"server"`
	Database    DatabaseConfig    `yaml:"database"`
	ExternalAPI ExternalAPIConfig `yaml:"external_api"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type ExternalAPIConfig struct {
	URL string `yaml:"url"`
}

var Cfg *Config

func LoadConfig() {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Failed to decode config file: %s", err)
	}

	Cfg = &config
}
