package engcore_modbus

import (
	"encoding/binary"
	"fmt"
)

type ModbusRawData struct {
	TransactionID   uint16 //Transaction ID
	ProtocolID      uint16 // 0000
	Length          uint16 // byte length  UnitID+FunctionalCodeData
	UnitID          uint8  //SlaveAdress
	FucnctionalCode uint8  //Functional Code
	Data            []byte //DATA
}

func RawDataToModbusRawData(data []byte) (*ModbusRawData, error) {
	frame := &ModbusRawData{
		TransactionID:   binary.BigEndian.Uint16(data[0:2]),
		ProtocolID:      binary.BigEndian.Uint16(data[2:4]),
		Length:          binary.BigEndian.Uint16(data[4:6]),
		UnitID:          data[6],
		FucnctionalCode: data[7],
		Data:            data[8:],
	}
	if int(frame.Length) != len(frame.Data)+2 {
		return nil, fmt.Errorf("Error ModbusTCp frame length")
	}
	return frame, nil
}
