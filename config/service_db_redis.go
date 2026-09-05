package config

type ServiceDBRedis struct {
	Name        string `yaml:"name"`         // 名称
	Addr        string `yaml:"addr"`         // 连接地址
	Password    string `yaml:"password"`     // 密码
	Port        uint16 `yaml:"port"`         // 端口号
	DB          int    `yaml:"db"`           // 数据库索引号
	Active      int    `yaml:"active"`       // 活跃连接数
	Idle        int    `yaml:"idle"`         // 空闲连接数
	IdleTimeout int    `yaml:"idle_timeout"` // 最大生命时长, 单位秒
}
