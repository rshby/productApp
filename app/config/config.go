package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppConfig *AppConfig
}

// function provider
func NewConfigApp() IConfig {
	cfg := viper.New()
	cfg.SetConfigFile("config.json")
	cfg.SetConfigType("json")
	cfg.AddConfigPath("./")

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatalf("cant load config : %v", err)
	}

	// create config
	config := &AppConfig{
		App: &App{
			Name:   cfg.GetString("app.name"),
			Author: cfg.GetString("app.author"),
			Port:   cfg.GetInt("app.port"),
		},
		Database: &Database{
			Host:     cfg.GetString("database.host"),
			User:     cfg.GetString("database.user"),
			Password: cfg.GetString("database.password"),
			Port:     cfg.GetInt("database.port"),
			Name:     cfg.GetString("database.name"),
		},
		Jaeger: &Jaeger{
			Host: cfg.GetString("jaeger.host"),
			Port: cfg.GetInt("jaeger.port"),
		},
		Logging: &Logging{
			Path: cfg.GetString("logging.path"),
		},
	}

	return &Config{AppConfig: config}
}

// method implementasi get config
func (c *Config) GetConfig() *AppConfig {
	return c.AppConfig
}

type AppConfig struct {
	App      *App      `json:"app,omitempty"`
	Database *Database `json:"database,omitempty"`
	Jaeger   *Jaeger   `json:"jaeger,omitempty"`
	Logging  *Logging  `json:"logging,omitempty"`
}

type App struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Port   int    `json:"port,omitempty"`
}

type Database struct {
	Host     string `json:"host,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Port     int    `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Logging struct {
	Path string `json:"path,omitempty"`
}

type Jaeger struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}
