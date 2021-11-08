package engcore_modbus

import "encoding/binary"

type ModbusRawData struct {
	TransactionID   uint16
	ProtocolID      uint16
	Length          uint16
	UnitID          uint8
	FucnctionalCode uint8
	Data            []byte
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
}
