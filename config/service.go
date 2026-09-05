package config

type Service struct {
	GroupName string           `yaml:"group_name"`
	Name      string           `yaml:"name"`
	Listen    *ServiceListen   `yaml:"listen"`
	Config    []*ServiceConfig `yaml:"config"`
	DB        *ServiceDB       `yaml:"db"`
	Rpc       []*ServiceRpc    `yaml:"rpc"`
	Log       *ServiceLog      `yaml:"log"`
}

type ServiceListen struct {
	Ip   string `yaml:"ip"`
	Port uint16 `yaml:"port"`
}

type ServiceLog struct {
	Path  string `yaml:"path"`
	Level uint8  `yaml:"level"`
}
