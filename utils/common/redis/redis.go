package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	// REDIS redis单例
	REDIS           *redis.Client
	LIMIT_LOGIN_KEY = "limit_login_%s_%s"
)

// Init 初始化
func Init() {
	fmt.Println("redis init")
	REDIS = InitRedis()
}

func InitRedis() *redis.Client {
	fmt.Println("redis.addr: ", viper.GetString("redis.addr"))
	options := redis.Options{Addr: viper.GetString("redis.addr"), DB: viper.GetInt("redis.db")}
	if viper.GetString("redis.password") != "" {
		options.Password = viper.GetString("redis.password")
	}
	client := redis.NewClient(&options)

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("redis connect fault, ", err)
		os.Exit(1)
	}
	return client
}

func Close() {
	if err := REDIS.Close(); err != nil {
		fmt.Println("redis close error, ", err)
	}
}
