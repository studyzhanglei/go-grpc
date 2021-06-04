package initialize

import (
	"content-grpc/global"
	"context"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
		PoolSize: 20,
		PoolTimeout:  30 * time.Second,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MinIdleConns: 5, //最小连接数量
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}

func RedisPool() {
	global.REDISPOOL = poolInitRedis(global.CONFIG.Redis.Addr, global.CONFIG.Redis.Password)
}

// redis pool
func poolInitRedis(server string, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     10,//空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   20,//最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
