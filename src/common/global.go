package common

import "github.com/gobuffalo/packr/v2"

// 全局配置
var (
    Conf    Configuration // 配置信息
    ConfBox *packr.Box    // 配置打包
)
