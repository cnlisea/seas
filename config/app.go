package config

type App struct {
	Config    *Config    `yaml:"config"`
	Discovery *Discovery `yaml:"discovery"`
	Services  []*Service `yaml:"services"`
}
