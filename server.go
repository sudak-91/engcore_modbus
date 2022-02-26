package engcore_modbus

import (
	"fmt"
	"log"
	"net"
)

type ModbusServer struct {
	ip     net.IP
	port   int
	mbMap  *ModbusRegisters
	action [17]func([]byte, *ModbusRegisters) ([]byte, error)
}

func NewModbusServer(ip string, port int, mMap *ModbusRegisters) *ModbusServer {

	m := &ModbusServer{}
	log.Println("create new modbus server instance")
	ipV4Adr, _, netparserr := net.ParseCIDR(ip)
	if netparserr != nil {
		log.Fatalln("invalid ip adress")
	}
	m.ip = ipV4Adr
	m.port = port
	m.mbMap = mMap
	m.action[1] = readCoilStatus
	m.action[2] = readInputStatus
	m.action[3] = readHoldingRegisters
	m.action[4] = readInputRegister
	m.action[5] = forseSingleCoil
	m.action[6] = presetSingleRegister
	m.action[15] = forseMultipalCoil
	m.action[16] = presetMultipalRegister
	log.Println("add modbus function to action list")
	return m
}

func (m *ModbusServer) StartServer() {

	ls, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: m.port,
		IP:   m.ip,
	})

	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("server was started. %v", ls.Addr())

	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Errorf("%v\n", err.Error())
		}

		//Client

		go m.handlerClient(conn)

	}

}

func (m *ModbusServer) handlerClient(conn net.Conn) {
	defer func() {
		log.Printf("Connection was closed")
		conn.Close()
		recover()
	}()

	log.Printf("Connected: %v", conn.RemoteAddr())
	buffer := make([]byte, 2048)

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Errorf("%v reading error %v", conn.RemoteAddr(), err.Error())
			continue
		}
		log.Printf("message length is: %v\n", length)
		bufferData := buffer[:length]
		frame, err := RawDataToModbusRawData(bufferData)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		fmt.Printf("Request frame:%v", frame)

		if frame.FunctionalCode > 17 {
			result, err := errorResponce(&frame)

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
		log.Printf("functional code is: %v", frame.FunctionalCode)
		result, err := m.action[frame.FunctionalCode](frame.Data, m.mbMap)
		log.Println("Result:", result)
		//Unlock
		if err != nil {
			log.Println(err.Error())
			frame.FunctionalCode = frame.FunctionalCode + 128
		}

		responceframe, err := createResponce(&frame, result)
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
