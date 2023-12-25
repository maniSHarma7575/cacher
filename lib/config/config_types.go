package config

import "time"

type Config struct {
	CacheConfig CacheConfig `json:"cache"`
}

type CacheConfig struct {
	Enabled     bool          `json:"enabled"`
	Prefix      string        `json:"prefix"`
	Type        string        `json:"type"`
	TTL         time.Duration `json:"ttl"`
	RedisConfig RedisConfig   `json:"redis"`
}

type RedisConfig struct {
	Cluster      bool          `json:"cluster"`
	Endpoints    []string      `json:"endpoints"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
	DialTimeout  time.Duration `json:"dialTimeout"`
	PoolSize     int           `json:"poolSize"`
	PoolTimeout  time.Duration `json:"poolTimeout"`
	MaxConnAge   time.Duration `json:"maxConnAge"`
	IdleTimeout  time.Duration `json:"idleTimeout"`
}
