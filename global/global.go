package global

import (
	"content-grpc/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)


var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VIPER     *viper.Viper
	LOG   *zap.Logger
	GRPC  *grpc.Server
)
