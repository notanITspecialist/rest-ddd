package config

type Postgresql struct {
	DSN string `mapstructure:"dsn"`
}
