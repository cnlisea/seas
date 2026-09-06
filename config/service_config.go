package config

type ServiceConfig struct {
	Key     string              `yaml:"key"`     // 别名
	Path    string              `yaml:"path"`    // 结构包所在路径
	Name    string              `yaml:"name"`    // 结构名称
	Channel int                 `yaml:"channel"` // 渠道 0-本地 1-nacos
	Local   *ServiceConfigLocal `yaml:"local"`   // local配置
	Nacos   *ServiceConfigNacos `yaml:"nacos"`   // nacos配置
}

type ServiceConfigLocal struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type ServiceConfigNacos struct {
	GroupId    string                     `yaml:"group_id"`
	DataId     string                     `yaml:"data_id"`
	UpdateHook []*ServiceConfigUpdateHook `yaml:"update_hook"` // 更新回调
}

type ServiceConfigUpdateHook struct {
	Path string `yaml:"path"` // 函数所在路径
	Name string `yaml:"name"` // 函数名称
}
