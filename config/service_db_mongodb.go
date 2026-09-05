package config

type ServiceDBMongoDB struct {
	Name        string   `yaml:"name"`
	User        string   `yaml:"user"`         // 用户名
	Password    string   `yaml:"password"`     // 密码
	Addr        []string `yaml:"addr"`         // 连接地址
	DBName      string   `yaml:"db_name"`      // 数据库名称
	ReplicaSet  string   `yaml:"replica_set"`  // 副本集
	Timeout     int      `yaml:"timeout"`      // 超时时间
	Active      int      `yaml:"active"`       // 活跃连接数
	Idle        int      `yaml:"idle"`         // 空闲连接数
	IdleTimeout int      `yaml:"idle_timeout"` // 最大生命时长, 单位秒
}
