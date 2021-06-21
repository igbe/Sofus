package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads a config file with name `fileName` and extention `ext`
// located in `path` and stores the retrieved values in `configValues` a struct.
func LoadConfig(path, fileName, ext string, configValues interface{}) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read config file %s.%s located in %s: %v", fileName, ext, path, err)
	}
	err = viper.Unmarshal(&configValues)
	if err != nil {
		return fmt.Errorf("could not unmarshal config values into provided struct: %v", err)
	}
	return nil
}
