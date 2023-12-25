package config

import (
	"encoding/json"
	"path/filepath"

	"github.com/spf13/viper"
)

func Load(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)

	ext := filepath.Ext(configPath)
	fileName := filepath.Base(configPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext[1:])
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	cacheConfig := viper.Get("cache")

	jsonbody, err := json.Marshal(cacheConfig)
	if err != nil {
		return nil, err
	}

	config := Config{}
	if err := json.Unmarshal(jsonbody, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
