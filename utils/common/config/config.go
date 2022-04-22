// Package config 改成配置中心设计
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
)

// ConfigName 配置信息汇总
type ConfigName struct {
	name string
}

// Config config params
type Config struct {
	DB struct {
		DSN    string `yaml:"dsn"`
		Driver string `yaml:"driver"`
	}

	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

var (
	configOnce sync.Once
	// FileConfig instance of file config
	FileConfig Config
)

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
		cfgFile = path + "/utils/common/config/config.yaml" // TODO 配置文件的设置
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

// ConfigurationContre 配置中心
func ConfigurationContre() {
	// TODO 远程获取配置中心的参数，启动服务
}
