package engcore_modbus

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func StartServer() {

	mbMap := NewModbusMap()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	ls, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: 5002,
		IP:   net.IPv4(192, 168, 56, 101),
	})

	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("server start", ls.Addr())

	for {

		fmt.Println("Default")
		conn, err := ls.Accept()
		if err != nil {
			log.Panic(err.Error())
		}
		//Client
		go handlerClient(conn, mbMap)

	}

}

func handlerClient(conn net.Conn, mbMap *ModbusMap) {
	defer conn.Close()

	fmt.Println(conn.RemoteAddr())
	buffer := make([]byte, 512)

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			break
		}

		buffer = buffer[:length]
		log.Println(buffer)
		frame, err := RawDataToModbusRawData(buffer)
		if err != nil {
			log.Println(err.Error())
			break
		}
		fmt.Println("Request frame:", *frame)

		result, err := mbMap.Action[frame.FunctionalCode](frame.Data)
		log.Println("Result:", result)
		if err != nil {
			log.Println(err.Error())
			break
		}

		responceframe, err := createResponce(frame, result)
		log.Println("Responceframe:", responceframe)
		butesresult, err := responceframe.ModbusFrametoByteSlice()

		log.Println(butesresult)
		lenRcv, WriteError := conn.Write(butesresult)
		if WriteError != nil {
			fmt.Errorf(WriteError.Error())
			continue
		}
		if lenRcv == 0 {
			log.Default().Output(0, "Wweeeeee")

		}
	}

}

func createResponce(request *ModbusRawData, data []byte) (*ModbusRawData, error) {
	responce := &ModbusRawData{}
	responce.TransactionID = request.TransactionID
	responce.ProtocolID = request.ProtocolID
	responce.Length = (uint16)(len(data) + 2)
	responce.UnitID = request.UnitID
	responce.FunctionalCode = request.FunctionalCode
	responce.Data = data
	return responce, nil
}
