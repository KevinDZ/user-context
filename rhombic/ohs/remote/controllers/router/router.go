package router

import (
	"fmt"
	"user_context/rhombic/ohs/remote/controllers/api"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// Init @host 192.168.0.109:8080
// @BasePath /api/v1
func Init(g *gin.Engine) {
	prefixPath := viper.GetString("prefix_path")
	fmt.Println("prefix path: ", prefixPath)
	apiGroup := g.Group(prefixPath)
	router(apiGroup)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}


func router(Router *gin.RouterGroup) {
	v1 := Router.Group("/v1")
	v1.POST("/register", api.RegisteredController)
	v1.POST("/login", api.LoginController)
	v1.POST("/logout", api.LogoutController)
}
