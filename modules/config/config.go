package config

import (
	"log/slog"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type httpConfig struct {
	ListenAddress string
}

type Config struct {
	HTTP httpConfig
}

func getViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile("config.yml")
	return v
}

func NewConfig() (*Config, error) {
	slog.Info("Loading configuration")

	v := getViper()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	err = v.Unmarshal(&config)
	return &config, err
}

var Module = fx.Options(fx.Provide(NewConfig))
