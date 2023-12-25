package cacher

import (
	"github.com/maniSHarma7575/cacher/lib/cache"
	"github.com/maniSHarma7575/cacher/lib/config"
)

func New(configPath string) (cache.CacheInterface, error) {
	conf, err := config.Load(configPath)

	if err != nil {
		return nil, err
	}

	if !conf.CacheConfig.Enabled {
		return nil, nil
	}

	return nil, nil
}
