package config

type Secrets struct {
	GrpcToken string `yaml:"grpc_token"`
	JwtSecret string `yaml:"jwt_secret"`
}
