package config

type ServiceFlag struct {
	Name    string  `yaml:"name"`
	Help    string  `yaml:"help"`
	Short   rune    `yaml:"short"`
	Require bool    `yaml:"require"`
	Default *string `yaml:"default"`
	Env     string  `yaml:"env"`
	Key     string  `yaml:"key"`
}
