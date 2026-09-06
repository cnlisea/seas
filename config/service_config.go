package config

type ServiceConfig struct {
	Key     string              `yaml:"key"`     // 别名
	Path    string              `yaml:"path"`    // 结构包所在路径
	Name    string              `yaml:"name"`    // 结构名称
	Channel int                 `yaml:"channel"` // 渠道 0-本地 1-nacos
	Local   *ServiceConfigLocal `yaml:"local"`   // local配置
	Nacos   *ServiceConfigNacos `yaml:"nacos"`   // nacos配置
	Hook    *ServiceConfigHook  `yaml:"hook"`
}

type ServiceConfigLocal struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type ServiceConfigHook struct {
	Get    []*ServiceConfigHookVal `yaml:"get"`
	Update []*ServiceConfigHookVal `yaml:"update"`
}

type ServiceConfigHookVal struct {
	Path string `yaml:"path"` // 函数所在路径
	Name string `yaml:"name"` // 函数名称
}

type ServiceConfigNacos struct {
	GroupId string `yaml:"group_id"`
	DataId  string `yaml:"data_id"`
}

type ServiceConfigUpdateHook struct {
	Path string `yaml:"path"` // 函数所在路径
	Name string `yaml:"name"` // 函数名称
}
