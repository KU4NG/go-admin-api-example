package common

import "go.uber.org/zap/zapcore"

// Configuration 配置总入口
type Configuration struct {
    Server ServerConfiguration `mapstructure:"server" json:"server"`
    Logs   LogsConfiguration   `mapstructure:"logs" json:"logs"`
}

// ServerConfiguration 系统配置项
type ServerConfiguration struct {
    Name       string `mapstructure:"name" json:"name"`
    Host       string `mapstructure:"host" json:"host"`
    Port       int    `mapstructure:"port" json:"port"`
    Mode       string `mapstructure:"mode" json:"mode"`
    ApiPrefix  string `mapstructure:"api-prefix" json:"apiPrefix"`
    ApiVersion string `mapstructure:"api-version" json:"apiVersion"`
}

// LogsConfiguration 日志配置项
type LogsConfiguration struct {
    Level      zapcore.Level `mapstructure:"level" json:"level"`
    Colorful   bool          `mapstructure:"colorful" json:"colorful"`
    Path       string        `mapstructure:"path" json:"path"`
    MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
    MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
    MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
    Compress   bool          `mapstructure:"compress" json:"compress"`
}
