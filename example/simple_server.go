package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sudak-91/engcore_modbus"
)

func main() {
	mbMap := engcore_modbus.NewModbusMap()
	ModBusServer := engcore_modbus.NewModbusServer("0.0.0.0/24", 5002, mbMap)
	go ModBusServer.StartServer()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
}
