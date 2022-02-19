package engcore_modbus

import (
	"encoding/binary"
	"fmt"
	"log"
)

type ModbusRawData struct {
	TransactionID  uint16 //Transaction ID
	ProtocolID     uint16 // 0000
	Length         uint16 // byte length  UnitID+FunctionalCodeData
	UnitID         uint8  //SlaveAdress
	FunctionalCode uint8  //Functional Code
	Data           []byte //DATA
}

func RawDataToModbusRawData(data []byte) (ModbusRawData, error) {
	var frame ModbusRawData
	if len(data) < 8 {
		return frame, fmt.Errorf("length error\n")
	}

	frame.TransactionID = binary.BigEndian.Uint16(data[0:2])
	frame.ProtocolID = binary.BigEndian.Uint16(data[2:4])
	frame.Length = binary.BigEndian.Uint16(data[4:6])
	frame.UnitID = data[6]
	frame.FunctionalCode = data[7]
	frame.Data = data[8:]

	log.Printf("modbus frame is: %v\n", frame)
	if int(frame.Length) != len(frame.Data)+2 {
		return frame, fmt.Errorf("error ModbusTCP frame length")
	}
	return frame, nil
}
func (m *ModbusRawData) ModbusFrametoByteSlice() ([]byte, error) {
	data := make([]byte, m.Length+6)
	binary.BigEndian.PutUint16(data[:2], m.TransactionID)
	binary.BigEndian.PutUint16(data[2:4], m.ProtocolID)
	binary.BigEndian.PutUint16(data[4:6], m.Length)

	data[6] = m.UnitID
	data[7] = m.FunctionalCode
	copy(data[8:], m.Data)
	return data, nil
}
