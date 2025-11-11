package config

import (
	"log"

	"github.com/spf13/viper"
)

// centralize config loading with support for .env

type Config struct {
	DBURL         string `mapstructure:"db_url"`
	ArgoCDBaseURL string `mapstructure:"argocd_base_url"`
	GithubToken   string `mapstructure:"github_token"`
	ServiceName   string
}

func Load(service string) *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.SetEnvPrefix(service)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found:%v", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	cfg.ServiceName = service
	return &cfg
}
