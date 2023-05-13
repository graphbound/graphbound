package main

import (
	"fmt"

	"github.com/graphbound/graphbound/examples/quotes-api/pkg/yeapi"
	"github.com/graphbound/graphbound/pkg/config"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnvironment config.AppEnvironment `default:"development" envconfig:"APP_ENV"`
	YeAPIURL       yeapi.ClientURL       `default:"https://api.kanye.rest" envconfig:"YE_API_URL"`
}

func ProvideConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("QUOTES_API", &config); err != nil {
		return nil, fmt.Errorf("ProvideConfig: can not read config: %w", err)
	}
	return &config, nil
}
