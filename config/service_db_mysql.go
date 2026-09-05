package config

type ServiceDBMySQL struct {
	Name        string `yaml:"name"`         // 名称
	User        string `yaml:"user"`         // 用户名
	Password    string `yaml:"password"`     // 密码
	Addr        string `yaml:"addr"`         // 连接地址
	Port        uint16 `yaml:"port"`         // 端口号
	DBName      string `yaml:"db_name"`      // 数据库名称
	Timeout     string `yaml:"timeout"`      // 超时时间
	ParseTime   bool   `yaml:"parse_time"`   // 是否解析时间类型
	Loc         string `yaml:"loc"`          // 时区
	Charset     string `yaml:"charset"`      // 字符编码
	Active      int    `yaml:"active"`       // 活跃连接数
	Idle        int    `yaml:"idle"`         // 空闲连接数
	IdleTimeout int    `yaml:"idle_timeout"` // 最大生命时长, 单位秒
}
