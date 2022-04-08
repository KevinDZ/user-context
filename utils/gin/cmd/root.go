package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"user_context/rhombic/ohs/remote/controllers/router"
	"user_context/utils/gin/config"
	"user_context/utils/gin/setting"
)

var cfgFile string
var serverPort int

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "realibox auth server",
	Long:  "realibox auth server use help get more info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("启动参数: ", args)
		runServer()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $CURRENT_DIR/config.yaml)")
	rootCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "port on which the server will listen")
}

// 初始化配置
func initConfig() {
	fmt.Println("init config...")
	config.Init(cfgFile)
}

func runServer() {
	fmt.Println("server start...")
	setting.Init()
	defer func() {
		setting.RedisClose()
	}()

	//TODO 基础设施 - nacos SDK
	//go sdk.MircoServiceInfrastructure()

	gin.SetMode(viper.GetString("runmode"))
	fmt.Println("runmode: ", viper.GetString("runmode"))
	// 路由设置
	g := gin.Default()
	router.Init(g)
	g.Run(fmt.Sprintf(":%d", serverPort))

}

// Execute rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
