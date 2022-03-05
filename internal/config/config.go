package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig
}

type DBConfig struct {
	Server   string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_SERVER"`
	Password string `mapstructure:"DB_USER"`
	Name     string `mapstructure:"DB_PASSWORD"`
}

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init(configsDir string) (*Config, error) {

	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}


	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigFile(".env")
	
	if err := viper.ReadInConfig(); err != nil {
		return err
	}


	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	return viper.UnmarshalExact(&cfg.DBConfig)
}