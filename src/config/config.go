package config

type Configurations struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host        string
	Port        int
	Environment string
}

var Configuration *Configurations
