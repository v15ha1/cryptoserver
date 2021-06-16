package data

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

//NewConfig returns new Config
func NewConfig() *Config {
	return &Config{}
}

//Config application configuration
type Config struct {
	//Logging Log Params
	Logging struct {
		Level string `yaml:"level" envconfig:"LOGGING_LEVEL"`
		Sink  string `yaml:"sink" envconfig:"LOGGING_SINK"`
	} `yaml:"logging"`

	//Server Server Params
	Server struct {
		ListenAddress string `yaml:"listenAddress" envconfig:"SERVER_LISTEN_ADDRESS"`
	} `yaml:"server"`

	// HitBTC WSS Ticker - https://api.hitbtc.com/#subscribe-to-ticker
	HitBTC struct {
		Endpoint string `yaml:"endpoint" envconfig:"HITBTC_ENDPOINT"`
	} `yaml:"hitbtc"`

}

//Init initialize
func (cfg *Config) Init(filePath string) error {
	err := processFile(filePath, cfg)
	if err != nil {
		return err
	}

	err = processEnv(cfg)
	if err != nil {
		return err
	}
	return nil
}

func processFile(filePath string, cfg *Config) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func processEnv(cfg *Config) error {
	err := envconfig.Process(os.Args[0], cfg)
	if err != nil {
		return err
	}
	return nil
}
