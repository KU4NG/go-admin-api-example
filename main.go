package main

import (
    "context"
    "fmt"
    "go-admin-api-example/src/common"
    "go-admin-api-example/src/initialize"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // 配置初始化
    initialize.Config()

    // 日志初始化
    initialize.Logger()

    // 路由初始化
    r := initialize.Router()

    // 配置服务优雅启动
    // 参考官方文档：https://gin-gonic.com/zh-cn/docs/examples/graceful-restart-or-stop/
    srv := &http.Server{
        Addr:    fmt.Sprintf("%s:%d", common.Conf.Server.Host, common.Conf.Server.Port),
        Handler: r,
    }

    // 启动服务
    go func() {
        err := srv.ListenAndServe()
        // 启动时候如果报错，并且错误不是关闭服务器，则打印日志并退出
        if err != nil && err != http.ErrServerClosed {
            common.Log.Error(fmt.Sprintf("服务启动失败，%s", err.Error()))
        }
    }()

    /*
       通过用户传递的信号实现优雅的退出，如 windows 的 ctrl + c，Linux 的 kill
       Linux kill 信号说明：
       kill：默认发送 syscall.SIGTERM 信号
       kill -2：发送 syscall.SIGINT 信号
       kill -9：发送 syscall.SIGKILL 信号，但是没法捕捉到，所以不建议使用
    */
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit // 等待信号传入

    // 当停止信号传入时，给程序 5 秒钟的处理时间，避免没有处理完请求给客户端报错
    common.Log.Warn("开始停止服务...")
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()
    err := srv.Shutdown(ctx)
    if err != nil {
        common.Log.Error(fmt.Sprintf("服务停止失败：%s", err.Error()))
    }
    common.Log.Info("服务停止完成")
}
