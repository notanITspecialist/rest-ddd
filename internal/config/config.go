package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	AppServer Server `mapstructure:"server"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	v := viper.New()

	v.SetDefault("server.port", 8000)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal config into struct: %w", err)
	}

	return cfg, nil
}
