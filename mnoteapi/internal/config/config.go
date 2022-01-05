package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	Cache cache.CacheConf
}
