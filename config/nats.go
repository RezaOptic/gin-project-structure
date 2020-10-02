package config

type NatsConfig struct {
	Addresses     []string `yaml:"addresses"`
	NatsClusterID string   `yaml:"nats_cluster_id"`
}
