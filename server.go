package engcore_modbus

import (
	"fmt"
	"log"
	"net"
)

type ModbusServer struct {
	ModbusRegisters ModbusMap
}

func NewModbusServer() *ModbusServer {
	s := &ModbusServer{}
	s.ModbusRegisters.Coil = make([]ModbusCoil, 65000)
	s.ModbusRegisters.DiscreteInput = make([]ModbusCoil, 65000)
	s.ModbusRegisters.HoldingRegister = make([]ModbusRegister, 65000)
	s.ModbusRegisters.InputRegister = make([]ModbusRegister, 65000)
	return s
}

func (s *ModbusServer) StartServer() {
	ls, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: 5002,
		IP:   net.IPv4(192, 168, 56, 101),
	})
	log.Println(ls.Addr())
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Panic(err.Error())
		}
		fmt.Println(conn.RemoteAddr())
		buffer := make([]byte, 512)
		length, err := conn.Read(buffer)
		if err != nil {
			log.Panic(err.Error())
		}
		buffer = buffer[:length]

	}
}
