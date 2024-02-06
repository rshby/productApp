package mock

import (
	"github.com/stretchr/testify/mock"
	"productApp/app/config"
)

type ConfigMock struct {
	Mock mock.Mock
}

// function provider
func NewConfigMock() *ConfigMock {
	return &ConfigMock{mock.Mock{}}
}

// method implement get config
func (c *ConfigMock) GetConfig() *config.AppConfig {
	args := c.Mock.Called()

	value := args.Get(0)
	if value == nil {
		return nil
	}

	return value.(*config.AppConfig)
}
