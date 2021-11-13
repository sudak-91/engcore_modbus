package engcore_modbus

import (
	"bytes"
	"encoding/gob"
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
		fmt.Println(frame)
		result, err := mbMap.Action[frame.FunctionalCode](frame.Data)
		if err != nil {
			continue
		}

		frame.Data = result
		buf := bytes.Buffer{}
		enc := gob.NewEncoder(&buf)
		err2 := enc.Encode(frame)
		if err2 != nil {
			fmt.Errorf(err2.Error())
			continue
		}
		lenRcv, WriteError := conn.Write(buf.Bytes())
		if WriteError != nil {
			fmt.Errorf(WriteError.Error())
			continue
		}
		if lenRcv == 0 {
			log.Default().Output(0, "Wweeeeee")

		}
	}
}
