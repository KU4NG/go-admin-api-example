package common

import (
    "github.com/gobuffalo/packr/v2"
    "go.uber.org/zap"
    "gorm.io/gorm"
)

// 全局配置
var (
    Conf    Configuration      // 配置信息
    ConfBox *packr.Box         // 配置打包
    Log     *zap.SugaredLogger // 日志输出
    DB      *gorm.DB           // 数据库
)

// 时间格式化常量
const (
    MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
    SecLocalTimeFormat  = "2006-01-02 15:04:05"
    DateLocalTimeFormat = "2006-01-02"
)
