package config

type IConfig interface {
	GetConfig() *AppConfig
}
