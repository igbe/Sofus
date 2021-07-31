package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration holds the urls of the various supported sites
type Configuration struct {
	Orgs    []string          `mapstructure:"urlConfig"`
	Workers map[string]string `mapstructure:"workerConfig"`
}

// LoadConfig loads a config file with name `fileName` and extention `ext`
// located in `path` and stores the retrieved values in `conf` a struct
// which it also returns.
func LoadConfig(path, fileName, ext string) (interface{}, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext)

	var conf *Configuration

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("could not read config file %s.%s located in %s: %v", fileName, ext, path, err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config values into provided struct: %v", err)
	}
	return *(conf), nil
}
