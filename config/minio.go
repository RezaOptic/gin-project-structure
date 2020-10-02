package config

type MinIOConfig struct {
	Addresses       string `yaml:"addresses"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
}
