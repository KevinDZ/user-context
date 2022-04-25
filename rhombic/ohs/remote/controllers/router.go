package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-context/rhombic/ohs/local/pl/errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"user-context/rhombic/ohs/local/pl"
	"user-context/rhombic/ohs/local/pl/vo"
	"user-context/rhombic/ohs/local/services"
	"user-context/utils/common"
)

func Init(g *gin.Engine) {
	prefixPath := viper.GetString("prefix_path")
	fmt.Println("prefix path: ", prefixPath)
	apiGroup := g.Group(prefixPath)
	router(apiGroup) // realihub-用户上下文
	open(apiGroup)   // TODO 开放平台-用户上下文
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func open(Router *gin.RouterGroup) {
	open := Router.Group("/open")
	open.POST("/index", func(context *gin.Context) {
		// TODO
	})
}

// realihub系统的用户上下文
func router(Router *gin.RouterGroup) {
	v1 := Router.Group("/v1")
	v1.POST("/register", registeredController)
	v1.POST("/login", loginController)
	v1.POST("/logout", logoutController)
}

// registeredController 注册控制器
func registeredController(ctx *gin.Context) {
	var request vo.RegisteredRequest
	request.Method = ctx.Request.Method
	request.IP = ctx.Request.RemoteAddr
	request.Proxies = ctx.Request.URL.RequestURI()
	// 0.解析参数
	if err := ctx.ShouldBind(request); err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 捕获异常
	defer errors.Recover("", request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, 500)

	request.RootID = "user" // TODO 设置聚合根ID
	// 1.注册用户，密码加密
	request.PassWord = common.PassWordEncryption(request.PassWord)
	res, err := services.RegisteredAppService(request)
	if err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 2.登录
	result, token, loginErr := services.LoginAppService(res)
	if loginErr != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, loginErr.Error(), 500)
		ctx.JSON(http.StatusBadRequest, loginErr)
		return
	}

	// 3.返回结果脱敏
	respond := vo.LoginRespond{
		Code:    pl.Success,
		Message: "",
		Info: vo.LoginInfo{
			ID:        result.UserID,
			Mobile:    result.Mobile,
			Email:     result.Email,
			AvatarUrl: result.AvatarUrl,
			AuthToken: token,
		},
	}
	respondJson, respondErr := json.Marshal(respond)
	if respondErr != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, respondErr.Error(), 500)
		ctx.JSON(http.StatusBadRequest, respondErr)
		return
	}
	ctx.JSON(http.StatusOK, respondJson)
}

// loginController 登录控制器
func loginController(ctx *gin.Context) {
	// 解析http通信获取的数据，转义成应用服务可识别的数据
	var request vo.LoginRequest
	request.Method = ctx.Request.Method
	request.IP = ctx.Request.RemoteAddr
	request.Proxies = ctx.Request.URL.RequestURI()
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Format{
			Level:      "",
			Now:        time.Now().Unix(),
			Method:     request.Method,
			ClientType: "",
			Client:     request.IP,
			Proxies:    request.Proxies,
			Message:    err.Error(),
			Code:       500,
		})
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 捕获异常
	defer errors.Recover("", request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, 500)

	request.RootID = "user" // TODO 设置聚合根ID
	result, token, err := services.LoginAppService(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.Format{
			Level:      "",
			Now:        time.Now().Unix(),
			Method:     request.Method,
			ClientType: "",
			Client:     request.IP,
			Proxies:    request.Proxies,
			Message:    err.Error(),
			Code:       500,
		})
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 返回结果脱敏
	respond := vo.LoginRespond{
		Code:    pl.Success,
		Message: "",
		Info: vo.LoginInfo{
			ID:        result.UserID,
			AuthToken: token,
		},
	}
	respondJson, respondErr := json.Marshal(respond)
	if respondErr != nil {
		ctx.JSON(http.StatusBadRequest, respondErr)
		return
	}
	ctx.JSON(http.StatusOK, respondJson)
}

// logoutController 登出控制器
func logoutController(ctx *gin.Context) {
	var request vo.LogoutRequest
	request.Method = ctx.Request.Method
	request.IP = ctx.Request.RemoteAddr
	request.Proxies = ctx.Request.URL.RequestURI()
	if err := ctx.ShouldBind(request); err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 捕获异常
	defer errors.Recover("", request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, 500)

	if err := services.LogoutAppService(request); err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(pl.Success, "")
}
