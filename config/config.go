package config

import (
	"log"
	"os"
	"sync"

	"github.com/gosidekick/goconfig"
)

// Config struct contains all configurations and parameters
// of the current program.
type Config struct {
	DiscountAddr string `json:"discount_addr" cfg:"discount_addr" cfgDefault:"localhost:5050"`
	Host         string `json:"host" cfg:"host" cfgDefault:"localhost"`
	Port         int    `json:"port" cfg:"port" cfgDefault:"8080"`
}

var (
	cfg  *Config
	once sync.Once
)

// Get returns an instance of the config settings
func Get() *Config {
	once.Do(func() {
		goconfig.PrefixEnv = "HASH"
		cfg = &Config{}
		err := goconfig.Parse(cfg)
		if err != nil {
			log.Println("error parsing config:", err)
			os.Exit(-1)
		}
	})
	return cfg
}
