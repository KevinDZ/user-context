package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config 配置信息汇总
type Config struct {
	name string
}

// Init 读取配置文件
func Init(cfgFile string) {
	t1 := os.Getenv("ENV_REDIS_ADDR")
	fmt.Println("env.. : ", t1)
	cluster := os.Getenv("ENV_CLUSTER_ADDR")
	fmt.Println("cluster env: ", cluster)
	namespace := os.Getenv("ENV_CLUSTER_NAMESPACE")
	fmt.Println("namespace env: ", namespace)
	if cfgFile == "" {
		path, _ := os.Getwd()
		cfgFile = path + "/config.yaml"
		fmt.Println(cfgFile)

	}

	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix("ENV")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
}
