package config

type HttpServer struct {
	Host string `koanf:"host" validate:"required"`
	Port int    `koanf:"port" validate:"required"`
}
