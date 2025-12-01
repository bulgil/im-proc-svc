package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env           string        `yaml:"env" env-default:"prod"`
	HTTPServerCfg HTTPServerCfg `yaml:"http_server_cfg" env-required:"true"`
	PGCfg         PGCfg         `yaml:"pg_cfg" env-required:"true"`
}

type HTTPServerCfg struct {
	Addr string `yaml:"addr" env-default:":9999"`
}

type PGCfg struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-default:"im-proc-svc"`
}

func ReadConfig() Config {
	var cfgPath string

	flag.StringVar(&cfgPath, "c", "./config/config.yaml", "path to the config")
	flag.Parse()

	if cfgPath == "" {
		panic("config path is not provided")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file doesn't exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic(fmt.Sprintf("error with read config: %s", err.Error()))
	}

	return cfg
}
