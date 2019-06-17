package config

import (
	"fmt"

	"github.com/paked/configure"
)

// Config includes all configurations for the app
type Config struct {
	MongodbServerHost *string
	MongodbPort       *string
	MongodbName       *string
	APIPort           *string
}

func newConfig() (*Config, *configure.Configure) {
	var conf = configure.New()
	var cfg = &Config{
		MongodbServerHost: conf.String("mongodbServerHost", "localhost", "The host address with which the mongodb instance containing the network setting database will be reachable"),
		MongodbPort:       conf.String("mongodbPort", "27017", "Active port for mongodb server"),
		MongodbName:       conf.String("mongodbName", "betta_fish_house", "The name of database being used."),
		APIPort:           conf.String("apiPort", "8888", "Port used for APIs, default to 8888 to work without root privileges"),
	}
	return cfg, conf
}

// NewConfig returns configurations parsed from json file or command line flag
func NewConfig(jsonPath string) *Config {
	cfg, conf := newConfig()
	conf.Use(configure.NewFlag())
	conf.Use(configure.NewJSONFromFile(jsonPath))
	conf.Use(configure.NewEnvironment())
	conf.Parse()
	return cfg
}

// Print all configurations
func (cfg *Config) Print() {
	fmt.Println("Parsed configurations:")

	fmt.Println("Mongo Server Host:", *cfg.MongodbServerHost)
	fmt.Println("Mongo Port:", *cfg.MongodbPort)
	fmt.Println("Database 's name:", *cfg.MongodbName)
	fmt.Println("API Port:", *cfg.APIPort)
}
