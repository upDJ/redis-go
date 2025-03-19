package config

import (
	"flag"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		flag.StringVar(&instance.Dir, "dir", "./", "Path to directory where the RDB file is stored")
		flag.StringVar(&instance.DBFilename, "dbfilename", "default", "Name of the RDB file")
		flag.Parse()
	})
	return instance
}

type Config struct {
	Dir        string
	DBFilename string
}
