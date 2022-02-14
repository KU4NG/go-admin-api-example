package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Cors 跨域配置中间件
func Cors(ctx *gin.Context) {
    method := ctx.Request.Method
    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
    ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
    ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
    ctx.Header("Access-Control-Allow-Credentials", "true")

    // 处理 OPTION 请求
    if method == "OPTIONS" {
        ctx.AbortWithStatus(http.StatusNoContent)
    }

    ctx.Next()
}
