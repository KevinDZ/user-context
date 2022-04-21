package api

import (
	"encoding/json"
	"net/http"
	"user-context/rhombic/ohs/local/pl"
	"user-context/rhombic/ohs/local/pl/vo"
	"user-context/rhombic/ohs/local/services"
	"user-context/utils/common"

	"github.com/gin-gonic/gin"
)

// RegisteredController 注册控制器
func RegisteredController(ctx *gin.Context) {
	// 1.注册
	var request vo.RegisteredRequest
	request.Method = ctx.Request.Method
	request.IP = ctx.Request.RemoteAddr
	request.Proxies = ctx.Request.URL.RequestURI()
	if err := ctx.ShouldBind(request); err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	request.RootID = "user" // TODO 设置聚合根ID
	// 注册用户，密码加密
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

// LoginController 登录控制器
func LoginController(ctx *gin.Context) {
	// 解析http通信获取的数据，转义成应用服务可识别的数据
	var request vo.LoginRequest
	request.Method = ctx.Request.Method
	request.IP = ctx.Request.RemoteAddr
	request.Proxies = ctx.Request.URL.RequestURI()
	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	request.RootID = "user" // TODO 设置聚合根ID
	result, token, err := services.LoginAppService(request)
	if err != nil {
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

// LogoutController 登出控制器
func LogoutController(ctx *gin.Context) {
	var request vo.LogoutRequest
	request.Method = ctx.Request.Method
	request.IP = ctx.Request.RemoteAddr
	request.Proxies = ctx.Request.URL.RequestURI()
	if err := ctx.ShouldBind(request); err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if err := services.LogoutAppService(request); err != nil {
		common.Errorln(request.Method, request.ClientType, request.SiteCode, request.IP, request.Proxies, err.Error(), 500)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(pl.Success, "")
}
