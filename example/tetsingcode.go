package main

import (
	"fmt"
	"internal/Mock"

	"github.com/sudak-91/engcore_modbus"
)

func main() {
	frame := Mock.GenerateWriteCoilsRequest(0, 10)
	fmt.Println(frame)
	frame2 := Mock.GenerateWriteRegisterRequest(0, 10)
	fmt.Println(frame2)
	k, _ := engcore_modbus.RawDataToModbusRawData(frame2)
	fmt.Println(k.TransactionID)
	fmt.Println(k.Length)

}
