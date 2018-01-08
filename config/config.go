package config

import (
	"fmt"
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

type Config struct {
	Polling PollingConfig
}

type PollingConfig struct {
	DB string
}

func NewConfig(configPaths ...string) (*Config, bool) {
	var cfg Config
	var ok bool

	env := os.Getenv("TKPENV")
	if env == "" {
		env = "development"
	}

	for _, configPath := range configPaths {
		configPath = fmt.Sprintf("%s/polling.%s.ini", configPath, env)
		fmt.Println(configPath)
		if err := gcfg.ReadFileInto(&cfg, configPath); err != nil {
			log.Println(err)
		} else {
			ok = true
			log.Printf("open %s succcessful", configPath)
			break
		}
	}

	return &cfg, ok
}
