package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sudak-91/engcore_modbus"
)

func main() {

	go engcore_modbus.StartServer()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
}
