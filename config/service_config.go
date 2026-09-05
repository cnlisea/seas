package config

type ServiceConfig struct {
	Key        string                   `yaml:"key"`         // 别名
	Path       string                   `yaml:"path"`        // 结构包所在路径
	Name       string                   `yaml:"name"`        // 结构名称
	Channel    int                      `yaml:"channel"`     // 渠道 0-本地 1-nacos
	Local      *ServiceConfigLocal      `yaml:"local"`       // local配置
	Nacos      *ServiceConfigNacos      `yaml:"nacos"`       // nacos配置
	UpdateHook *ServiceConfigUpdateHook `yaml:"update_hook"` // 更新回调
}

type ServiceConfigLocal struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type ServiceConfigNacos struct {
	GroupId string `yaml:"group_id"`
	DataId  string `yaml:"data_id"`
}

type ServiceConfigUpdateHook struct {
	MethodName string `yaml:"method_name"` // 自定义方法名称
}
