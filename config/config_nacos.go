package config

type ConfigNacos struct {
	NamespaceId string             `yaml:"namespace_id"` // 命名空间
	Nodes       []*ConfigNacosNode `yaml:"nodes"`        // 节点
}

type ConfigNacosNode struct {
	Addr string `yaml:"addr"` // 地址, 域名或IP
	Port uint16 `yaml:"port"` // 端口号
}
