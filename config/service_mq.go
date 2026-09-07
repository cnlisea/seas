package config

type ServiceMQ struct {
	NameServer []string              `yaml:"name_server"` // 名称服务地址
	AccessKey  string                `yaml:"access_key"`  // 访问Key
	SecretKey  string                `yaml:"secret_key"`  // 密钥key
	Namespace  string                `yaml:"namespace"`   // 名称空间
	Producer   []*ServiceMQProducer  `yaml:"producer"`    // 生产
	Subscribe  []*ServiceMQSubscribe `yaml:"subscribe"`   // 订阅
}

type ServiceMQProducer struct {
	Name    string `yaml:"name"`     // 名称
	GroupId string `yaml:"group_id"` // 组编号
}

type ServiceMQSubscribe struct {
	Name      string                      `yaml:"name"`       // 名称
	GroupId   string                      `yaml:"group_id"`   // 组编号
	Broadcast bool                        `yaml:"broadcast"`  // 广播模式
	BatchSize int                         `yaml:"batch_size"` // 批量
	Config    []*ServiceMQSubscribeConfig `yaml:"config"`     // 配置
}

type ServiceMQSubscribeConfig struct {
	Topic string   `yaml:"topic"` // 主题名称
	Tag   string   `yaml:"tag"`   // 标签
	Path  string   `yaml:"path"`  // 服务所在路径
	Name  string   `yaml:"name"`  // 服务名称
	Param []string `yaml:"param"` // 服务参数
}
