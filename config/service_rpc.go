package config

type ServiceRpc struct {
	Offline *bool    `yaml:"offline"` // 软状态离线
	Path    string   `yaml:"path"`    // 服务所在路径
	Name    string   `yaml:"name"`    // 服务名称
	Param   []string `yaml:"param"`   // 服务参数
}
