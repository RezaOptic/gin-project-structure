package service

import (
	"github.com/RezaOptic/gin-project-structure/config"
	"gitlab.com/RezaOptic/go-utils/logger"
	"gitlab.com/RezaOptic/go-utils/standriver"

	"os"
	"strings"
	"time"
)

var Sd *standriver.Connection

func NatsInit() {
	hostname, err := os.Hostname()
	if err != nil {
		logger.ZSLogger.Errorf("failed to get HostName with error: %v\n", err)
		os.Exit(-1)
	}
	hostname = strings.Replace(hostname, ".", "_", -1) + config.Configs.Service.ServiceName
	for {
		Sd = standriver.New(config.Configs.Nats.Addresses,
			config.Configs.Nats.NatsClusterID, hostname)
		err = Sd.Connect()
		if err != nil {
			logger.ZSLogger.Errorf("failed to connect to stan: %v\n", err)
			time.Sleep(time.Second*2)
		}
		return
	}

}
