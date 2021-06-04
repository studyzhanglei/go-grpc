package global

import (
	"content-grpc/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	redigo "github.com/gomodule/redigo/redis"
)


var (
	DB     *gorm.DB
	REDISPOOL  *redigo.Pool
	REDIS *redis.Client
	CONFIG config.Server
	VIPER     *viper.Viper
	LOG   *zap.Logger
	GRPC  *grpc.Server
)
