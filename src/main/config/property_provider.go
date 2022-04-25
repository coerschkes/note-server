package config

import (
	"github.com/magiconair/properties"
)

type ConfigProvider struct {
	properties *properties.Properties
}

func MakeConfigProvider() ConfigProvider {
	return ConfigProvider{properties: properties.MustLoadFile("resources/server.properties", properties.UTF8)}
}

func (p ConfigProvider) GetProperty(propertyKey string) string {
	return p.properties.MustGetString(propertyKey)
}
