package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
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
