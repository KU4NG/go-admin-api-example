package initialize

import (
    "bytes"
    "fmt"
    "github.com/gobuffalo/packr/v2"
    "github.com/spf13/viper"
    "go-admin-api-example/src/common"
    "log"
    "os"
    "strings"
)

// 配置文件相关信息
const (
    EnvName        = "RUN_ENV"                 // 环境变量名称，用于判定运行环境，以此选择默认配置文件
    ConfigBoxName  = "go-admin-api-config-box" // packr 用于存放配置的命名空间
    ConfigType     = "yml"                     // 配置文件类型
    ConfigPath     = "../config"               // 配置文件目录
    ConfigDevFile  = "application.dev.yml"     // 开发环境配置文件名称
    ConfigProdFile = "application.prod.yml"    // 生产环境配置文件名称
)

// ReadConfig 读取配置文件方法封装
func ReadConfig(v *viper.Viper, file string) {
    // 属性设置
    v.SetConfigType(ConfigType)

    // 判断配置是否存在
    config, err := common.ConfBox.Find(file)
    if err != nil {
        panic(fmt.Sprintf("配置文件读取失败：%s", err.Error()))
    }

    // 加载配置文件
    err = v.ReadConfig(bytes.NewReader(config))
    if err != nil {
        panic(fmt.Sprintf("配置文件加载失败：%s", err.Error()))
    }
}

// Config 配置文件初始化
func Config() {
    // 设置打包配置
    common.ConfBox = packr.New(ConfigBoxName, ConfigPath)

    // 读取配置文件，默认加载 dev 环境所有配置，所以 dev 环境配置必须包含所有配置项
    v := viper.New()
    ReadConfig(v, ConfigDevFile)

    // 设置配置
    settings := v.AllSettings()
    for i, setting := range settings {
        v.SetDefault(i, setting)
    }

    // 获取系统运行环境
    runEnv := strings.ToLower(os.Getenv(EnvName))

    // 根据不同的运行环境在追加修改配置项
    configFile := ""
    if runEnv == "prod" {
        configFile = ConfigProdFile
    }

    // 如果不是 dev 环境，则继续加载其它配置
    if configFile != "" {
        ReadConfig(v, configFile)
    }

    // 将配置转换成结构体，便于程序使用
    err := v.Unmarshal(&common.Conf)
    if err != nil {
        panic(fmt.Sprintf("配置文件解析失败：%s", err.Error()))
    }

    // 打印信息
    log.Println("配置文件初始化完成")
}
