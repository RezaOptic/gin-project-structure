package main

import (
	"flag"
	"github.com/RezaOptic/gin-project-structure/config"
	"github.com/RezaOptic/gin-project-structure/repository"
	grpcEngine "github.com/RezaOptic/gin-project-structure/router/grpc"
	httpEngine "github.com/RezaOptic/gin-project-structure/router/http"
	"github.com/RezaOptic/gin-project-structure/service"
	"os"
	"os/signal"
	"sync"
	"syscall"
)


func main() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")

	// init service configs
	config.Init(configFlag, prodConfigPath)

	// init repositories
	repository.Init()
	service.NatsInit()

	// run http and grpc servers
	wg := sync.WaitGroup{}
	wg.Add(2)
	go httpEngine.Run(config.Configs.Service.HttpPort)
	go grpcEngine.Run(config.Configs.Service.GrpcPort)

	// handle os signals
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		select {
		case <-sigc:
			// wg.Done()
			// TODO...
			os.Exit(1)
		}
	}()

	wg.Wait()
}
