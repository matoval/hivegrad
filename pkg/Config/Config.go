package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	IsHub bool
	Path string
	Name string
	Type string
}

func NewConfig() *Config {
	c := &Config{
		IsHub: false,
		Path: "./",
		Name: "test",
		Type: "yaml",
	}
	return c
}

func (c *Config) SetConfigPath(p string) *Config {
	c.Path = p
	return c
}

func (c *Config) SetConfigName(n string) *Config {
	c.Name = n
	return c
}

func (c *Config) SetConfigType(t string) *Config {
	c.Type = t
	return c
}

func (c *Config) LoadConfig() {
	viper.AddConfigPath(c.Path)
	viper.SetConfigName(c.Name)
	viper.SetConfigType(c.Type)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
