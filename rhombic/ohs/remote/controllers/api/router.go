package api

import (
	"encoding/json"
	"net/http"
	"user_context/rhombic/ohs/local/pl"
	"user_context/rhombic/ohs/local/services"

	"github.com/gin-gonic/gin"
)

//
func RegisteredController(ctx *gin.Context)  {
	var request pl.RegisteredRequest
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	services.RegisteredAppService(request)
}

// LoginController 登录控制器
func LoginController(ctx *gin.Context) {
	// 解析http通信获取的数据，转义成应用服务可识别的数据
	var request pl.LoginRequest
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
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
	ctx.JSON(http.StatusOK, resultJson)
}

// LogoutController 登出控制器
func LogoutController(ctx *gin.Context)  {
	var request pl.LogoutRequest
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if err := services.LogoutAppService(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "")
}