package config

type ScylladbConfig struct {
	Addresses []string `yaml:"addresses"`
	KeySpace  string   `yaml:"key_space"`
}
