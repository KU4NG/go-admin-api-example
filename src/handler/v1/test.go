package v1

import (
    "github.com/gin-gonic/gin"
    "go-admin-api-example/src/common"
    "net/http"
)

// PingHandler 测试接口处理函数
func PingHandler(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
        "name":    common.Conf.Server.Name,
        "message": "pong",
        "version": "build.2022.02.14.161700",
    })
}
