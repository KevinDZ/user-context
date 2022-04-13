package api

import (
	"encoding/json"
	"net/http"
	"user-context/rhombic/ohs/local/pl"
	"user-context/rhombic/ohs/local/pl/vo"
	"user-context/rhombic/ohs/local/services"

	"github.com/gin-gonic/gin"
)

// RegisteredController 注册控制器
func RegisteredController(ctx *gin.Context) {
	// 1.注册
	var request vo.RegisteredRequest
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	request.RootID = "user" // TODO 设置聚合根ID
	res, err := services.RegisteredAppService(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 2.登录
	result, err := services.LoginAppService(res)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(pl.Success, resultJson)
}

// LoginController 登录控制器
func LoginController(ctx *gin.Context) {
	// 解析http通信获取的数据，转义成应用服务可识别的数据
	var request vo.LoginRequest
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	request.RootID = "user" // TODO 设置聚合根ID
	result, err := services.LoginAppService(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(pl.Success, resultJson)
}

// LogoutController 登出控制器
func LogoutController(ctx *gin.Context) {
	var request vo.LogoutRequest
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if err := services.LogoutAppService(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(pl.Success, "")
}
