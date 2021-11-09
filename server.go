package engcore_modbus

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	mbMap := NewModbusMap()
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
		frame, err := RawDataToModbusRawData(buffer)
		if err != nil {
			continue
		}

	}
}
