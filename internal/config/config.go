package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	AppServer  Server     `mapstructure:"server"`
	Postgresql Postgresql `mapstructure:"postgresql"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	v := viper.New()

	v.SetDefault("server.port", 8000)

	v.SetDefault("postgresql.dsn", "postgres://root:password@127.0.0.1:5432/rest-ddd")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal config into struct: %w", err)
	}

	return cfg, nil
}
