package config

type Microservices struct {
	UserService microserviceInfo `yaml:"user_service"`
}

type microserviceInfo struct {
	Host     string `yaml:"host"`
	HttpPort string `yaml:"http_port"`
	GrpcPort string `yaml:"grpc_port"`
}
