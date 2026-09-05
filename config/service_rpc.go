package config

type ServiceRpc struct {
	Offline *bool    `yaml:"offline"` // 软状态离线
	Path    string   `yaml:"path"`    // 服务所在路径
	Name    string   `yaml:"name"`    // 服务名称
	Param   []string `yaml:"param"`   // 服务参数
}

type ServiceRpcParam struct {
	DBMySQL   *ServiceRpcParamType `yaml:"db_mysql"`
	DBRedis   *ServiceRpcParamType `yaml:"db_redis"`
	DBMongo   *ServiceRpcParamType `yaml:"db_mongo"`
	RpcClient *ServiceRpcParamType `yaml:"rpc_client"`
	Cfg       *ServiceRpcParamType `yaml:"cfg"`
	MQ        *ServiceRpcParamType `yaml:"mq"`
}

type ServiceRpcParamType struct {
	Index int `yaml:"index"` // 位置
	Value int `yaml:"value"` // 值
}
