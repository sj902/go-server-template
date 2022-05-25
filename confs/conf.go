package confs

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort string
}

var initialized bool
var conf Config

func Conf() *Config {
	if !initialized {
		log.Panic("conf is not initialized.")
	}
	return &conf
}

func Init(serverPort string) {
	conf.ServerPort = (serverPort)
	initialized = true
}
