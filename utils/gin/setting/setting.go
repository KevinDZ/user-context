package setting

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	// REDIS redis单例
	REDIS           *redis.Client
	LIMIT_LOGIN_KEY = "limit_login_%s_%s"
)

// Init 初始化
func Init() {
	fmt.Println("setting init")
	REDIS = InitRedis()
}
