package config

import (
	"flag"
	"os"
)
type HTTPServer struct{
     Addr   string
}

// env_default:"production"

type Config struct{
	Env           string `yaml:"env" env:"ENV" env-required:"true" `
	StoragePath   string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad(){
	var configPath string

	configPath=os.Getenv("CONFIG_PATH")

	if configPath==""{
		flags := flag.String("config","","path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath==""{
			log.fatal("Config path is not set")
		}
	}
}