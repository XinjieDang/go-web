package global

import (
	"com.dxj/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *config.Config // 全局Config
	DB     *gorm.DB
	Redis  *redis.Client
)
