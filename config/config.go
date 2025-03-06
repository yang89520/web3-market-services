package config

import (
	"github.com/urfave/cli/v2"

	"github.com/the-web3/market-services/flags"
)

type ServerConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type Config struct {
	Migrations string
	RpcServer  ServerConfig
	RestServer ServerConfig
	Metrics    ServerConfig
	MasterDB   DBConfig
	SlaveDB    DBConfig
}

func NewConfig(ctx *cli.Context) *Config {
	return &Config{
		Migrations: ctx.String("migrations"),
	}
}
