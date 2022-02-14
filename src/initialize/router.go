package initialize

import (
    "github.com/gin-gonic/gin"
    "go-admin-api-example/src/common"
    "go-admin-api-example/src/middleware"
    "go-admin-api-example/src/routes"
    "log"
)

// Router 路由初始化
func Router() *gin.Engine {
    // 设置运行模式
    gin.SetMode(common.Conf.Server.Mode)

    // 创建一个干净的路由引擎
    r := gin.New()

    // 中间件
    r.Use(middleware.AccessLog) // 访问日志中间件
    r.Use(middleware.Cors)      // 跨域中间件

    // 创建路由组，增加 /api 前缀
    apiGroup := r.Group(common.Conf.Server.ApiPrefix)

    // 创建路由组，增加 /api/v1 前缀
    v1Group := apiGroup.Group(common.Conf.Server.ApiVersion)
    {
        routes.TestRouters(v1Group) // 用于测试的路由
    }

    // 打印提示
    log.Println("路由初始化完成")

    // 返回路由
    return r
}
