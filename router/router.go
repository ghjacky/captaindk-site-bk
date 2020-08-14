package router

import "github.com/gin-gonic/gin"

type SHttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// api前缀定义
const apiPrefix = "/api/v1"

// 全局路由定义
var Router = gin.Default()

// 路由注册初始化
func RegisterRouter() {
	articleRouter()
	categoryRouter()
	FileRouter()
}

func newResponse(code int, message string, data interface{}) SHttpResponse {
	return SHttpResponse{Code: code, Message: message, Data: data}
}
