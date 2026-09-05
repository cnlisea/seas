package config

type ServiceDB struct {
	MySQL   []*ServiceDBMySQL   `yaml:"mysql"`
	Redis   []*ServiceDBRedis   `yaml:"redis"`
	MongoDB []*ServiceDBMongoDB `yaml:"mongodb"`
}
