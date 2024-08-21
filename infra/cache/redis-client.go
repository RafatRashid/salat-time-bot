package cache

import (
	"fmt"
	"log/slog"

	"github.com/go-redis/redis"
)

var client *redis.Client

type config struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func Connect() {
	conf := config{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		Db:       1,
	}

	slog.Info(fmt.Sprintf("connecting to redis at %s:%d...", conf.Host, conf.Port))

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
	})

	if _, err := client.Ping().Result(); err != nil {
		slog.Error(fmt.Sprintf("redis client error: %s", err.Error()))
		panic(err)
	}

	slog.Info(fmt.Sprintf("redis client connected on %s:%d", conf.Host, conf.Port))
}

func GetClient() *redis.Client {
	return client
}
