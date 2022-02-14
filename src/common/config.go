package common

// Configuration 配置总入口
type Configuration struct {
    Server ServerConfiguration `mapstructure:"server" json:"server"`
}

// ServerConfiguration 系统配置项
type ServerConfiguration struct {
    Name string `mapstructure:"name" json:"name"`
    Host string `mapstructure:"host" json:"host"`
    Port int    `mapstructure:"port" json:"port"`
    Mode string `mapstructure:"mode" json:"mode"`
}
