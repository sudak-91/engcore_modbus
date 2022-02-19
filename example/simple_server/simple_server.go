package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sudak-91/engcore_modbus"
)

func main() {
	mbMap, err := engcore_modbus.NewModbusRegisters(20, 20, 20, 20)
	if err != nil {
		panic(err)
	}
	mbMap.SetHoldingRegister(0, []uint16{12, 45, 67, 33445, 63, 1234, 12, 34})
	mbMap.SetInputRegister(0, []uint16{2345})
	mbMap.SetInputRegister(10, []uint16{12, 13, 14})
	ModBusServer := engcore_modbus.NewModbusServer("0.0.0.0/24", 5002, mbMap)
	go ModBusServer.StartServer()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
}
