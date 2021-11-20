package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sudak-91/engcore_modbus"
)

func main() {

	mbMap := engcore_modbus.NewModbusMap()
	go engcore_modbus.StartServer("0.0.0.0/24", 5002, mbMap)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
}
