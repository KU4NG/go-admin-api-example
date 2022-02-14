package routes

import (
    "github.com/gin-gonic/gin"
    v1 "go-admin-api-example/src/handler/v1"
)

// TestRouters 用于测试的路由组
func TestRouters(r *gin.RouterGroup) gin.IRoutes {
    r.GET("ping", v1.PingHandler)
    return r
}
