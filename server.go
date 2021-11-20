package engcore_modbus

import (
	"fmt"
	"log"
	"net"
)

func StartServer(ip string, port int, mbMap *ModbusMap) {
	ipv4Adr, _, netparserr := net.ParseCIDR(ip)
	if netparserr != nil {
		log.Fatalln("invalid ip adress")
	}

	ls, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: port,
		IP:   ipv4Adr,
	})

	if err != nil {
		log.Println(err.Error())
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
			continue
		}
		fmt.Println("Request frame:", *frame)

		if frame.FunctionalCode > 17 {
			result, err := errorResponce(frame)

			if err != nil {
				log.Println(err.Error())
				continue
			}

			lenrcv, err := conn.Write(result)
			if err != nil || lenrcv <= 0 {
				log.Println("error send recive with illegal functional code error", err.Error())
				continue
			}
			continue
		}

		result, err := mbMap.Action[frame.FunctionalCode](frame.Data)
		log.Println("Result:", result)
		if err != nil {
			log.Println(err.Error())
			frame.FunctionalCode = frame.FunctionalCode + 128
		}

		responceframe, err := createResponce(frame, result)
		log.Println("Responce Frame:", responceframe)
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

func errorResponce(request *ModbusRawData) ([]byte, error) {
	request.FunctionalCode = request.FunctionalCode + 128
	responce, err := createResponce(request, []byte{ILLEGAL_FUNCTION})
	if err != nil {
		log.Println("create pesponce err")
		return nil, err
	}
	log.Println("Responceframe:", responce)
	result, err := responce.ModbusFrametoByteSlice()
	if err != nil {
		log.Println("error convert modbus framme to byte slice")
		return nil, err
	}
	log.Println("Responce byte slice:", result)
	return result, nil

}
