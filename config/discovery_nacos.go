package config

type DiscoveryNacos struct {
	NamespaceId string                `yaml:"namespace_id"` // 命名空间
	Nodes       []*DiscoveryNacosNode `yaml:"nodes"`        // 节点
}

type DiscoveryNacosNode struct {
	Addr string `yaml:"addr"` // 地址, 域名或IP
	Port uint16 `yaml:"port"` // 端口号
}
