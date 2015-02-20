package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	conf "github.com/Bo0mer/enslaver/config"
	"github.com/Bo0mer/enslaver/enslaver"
	"github.com/Bo0mer/enslaver/facade"
	"github.com/Bo0mer/enslaver/slaveclient"

	"github.com/Bo0mer/os-agent/server"

	l4g "code.google.com/p/log4go"
)

func main() {
	config := loadConfig()

	slaveClient := slaveclient.NewSlaveClient()
	enslaver := enslaver.NewEnslaver(slaveClient)
	enslaverFacade := facade.NewEnslaverFacade(enslaver)

	registerSlaveHandler := server.NewHandler("POST", "/register", enslaverFacade.RegisterSlave)
	executeHandler := server.NewHandler("POST", "/jobs", enslaverFacade.CreateJob)

	s := server.NewServer(config.Server.Host, config.Server.Port)
	s.Register(registerSlaveHandler)
	s.Register(executeHandler)

	l4g.Info("Starting HTTP server on %s:%d", config.Server.Host, config.Server.Port)
	err := s.Start()
	if err != nil {
		l4g.Error("Unable to start server", err)
		return
	}
	l4g.Info("Start successful.")

	osChan := make(chan os.Signal)
	signal.Notify(osChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)

	<-osChan

	l4g.Info("Shtutting down HTTP server...")
	s.Stop()
	l4g.Info("Shutdown successsful.")
}

func loadConfig() (config conf.EnslaverConfig) {
	configDir := os.Getenv("ENSLAVER_CONFIG_DIR")
	configFile := fmt.Sprintf("%s%s", configDir, "/config.yml")
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		l4g.Error("Could not load configuration. Error: %s", err)
		panic("Could not load configuration!")
	}
	return config
}
