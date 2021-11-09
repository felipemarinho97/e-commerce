package config

import (
	"os"
	"sync"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/gosidekick/goconfig"
)

// Config struct contains all configurations and parameters
// of the current program.
type Config struct {
	DiscountAddr     string `json:"discount_addr" cfg:"discount_addr" cfgDefault:"localhost:5050"`
	Host             string `json:"host" cfg:"host" cfgDefault:"localhost"`
	Port             int    `json:"port" cfg:"port" cfgDefault:"8080"`
	ProductsMockFile string `json:"products_mock_file" cfg:"products_mock_file" cfgDefault:"./products.json"`
	IsBlackFriday    bool   `json:"is_black_friday" cfg:"is_black_friday" cfgDefault:"false"`
}

var (
	cfg  *Config
	once sync.Once
)

// Get returns an instance of the config settings.
func Get() *Config {
	once.Do(func() {
		goconfig.PrefixEnv = "HASH"
		cfg = &Config{}
		err := goconfig.Parse(cfg)
		if err != nil {
			common.Logger.LogError("error parsing config", err.Error())
			os.Exit(-1)
		}
	})

	return cfg
}
