package router

import "github.com/gin-gonic/gin"

// api前缀定义
const apiPrefix = "/api/v1"

// 全局路由定义
var Router = gin.Default()

// 路由注册初始化
func RegisterRouter() {
	articleRouter()
}
