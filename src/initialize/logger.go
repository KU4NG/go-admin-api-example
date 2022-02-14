package initialize

import (
    "fmt"
    "github.com/natefinch/lumberjack"
    "go-admin-api-example/src/common"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "log"
    "os"
    "time"
)

// ZapLocalTimeEncoder 定义日志时间格式
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format(common.MsecLocalTimeFormat))
}

// Logger 日志初始化
func Logger() {
    // 定义基础变量
    now := time.Now()                                                                                                                       // 当前时间
    filename := fmt.Sprintf("%s/%s.%04d-%02d-%02d.log", common.Conf.Logs.Path, common.Conf.Server.Name, now.Year(), now.Month(), now.Day()) // 日志保存路径
    hook := &lumberjack.Logger{
        Filename:   filename,                    // 日志保存路径
        MaxSize:    common.Conf.Logs.MaxSize,    // 文件最大
        MaxAge:     common.Conf.Logs.MaxAge,     // 文件保留天数
        MaxBackups: common.Conf.Logs.MaxBackups, // 文件备份个数
        Compress:   common.Conf.Logs.Compress,   // 是否压缩
    }
    defer hook.Close()

    // 简易配置 zap
    enConfig := zap.NewProductionEncoderConfig()
    enConfig.EncodeTime = ZapLocalTimeEncoder

    // 日志等级颜色输出处理
    if common.Conf.Logs.Colorful {
        enConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    } else {
        enConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    }

    // 日志输出位置配置(控制台/文件输出)
    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(enConfig),
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
        common.Conf.Logs.Level,
    )

    // 处理日志输出中打印当前文件的问题
    logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

    // 配置全局答应变量
    common.Log = logger.Sugar()

    // 打印输出日志
    log.Println("日志初始化完成")
}
