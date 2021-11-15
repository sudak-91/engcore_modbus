package engcore_modbus

import (
	"encoding/binary"
	"fmt"
	"log"
)

//Holding and Input Register
type ModbusRegister struct {
	Description string
	Value       uint16
}

//discrete Input and Coil
type ModbusCoil struct {
	Description string
	Value       byte
}

type ModbusMap struct {
	DiscreteInput   []ModbusCoil
	Coil            []ModbusCoil
	InputRegister   []ModbusRegister
	HoldingRegister []ModbusRegister
	Action          [17]func([]byte) ([]byte, error)
}

func NewModbusMap() *ModbusMap {
	m := &ModbusMap{}
	m.Coil = make([]ModbusCoil, 65535)
	m.DiscreteInput = make([]ModbusCoil, 65535)
	m.InputRegister = make([]ModbusRegister, 65535)
	m.HoldingRegister = make([]ModbusRegister, 65535)
	m.Action[1] = m.readCoilStatus
	m.Action[2] = m.ReadInputStatus
	m.Action[3] = m.ReadHoldingRegisters
	m.Action[4] = m.ReadInputRegister
	m.Action[5] = m.ForseSingleCoil
	m.Action[6] = m.PresetSingleRegister
	m.Action[15] = m.ForseMultipalCoil
	m.Action[16] = m.PresetMultipalRegister
	return m
}

//Return: data, error
//Mosbuc Command 0x01
func (m *ModbusMap) readCoilStatus(data []byte) ([]byte, error) {
	log.Println("Read Coil")
	offset := GetOffset(data)
	length := GetLength(data)
	if length > 65535 {
		return nil, fmt.Errorf("max coil count is 65535")
	}

	resultLength := length / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1) // slice for length and result
	Result[0] = byte(resultLength)

	for k, value := range m.Coil[offset : offset+length] {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = byte(1 << shift)
		}

	}
	return Result, nil

}

//Modbus 0x02
func (m *ModbusMap) ReadInputStatus(data []byte) ([]byte, error) {
	log.Println("Read Input Status")
	offset := GetOffset(data)
	length := GetLength(data)
	if length > 65535 {
		return nil, fmt.Errorf("max coil count is 65535")
	}

	resultLength := length / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1) // slice for length and result
	Result[0] = byte(resultLength)

	for k, value := range m.InputRegister[offset : offset+length] {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = byte(1 << shift)
		}

	}
	return Result, nil
}

//Modbus 0x03
func (m *ModbusMap) ReadHoldingRegisters(data []byte) ([]byte, error) {
	log.Println("Read Holding Register")
	offset := GetOffset(data)
	length := GetLength(data)
	if length > 65535 {
		return nil, fmt.Errorf("out of slice size")
	}
	byteCount := length * 2
	Result := make([]byte, byteCount+1)
	Result[0] = byte(byteCount)
	for i, value := range m.HoldingRegister[offset : length+offset] {
		binary.BigEndian.PutUint16(Result[i*2:(i+1)*2], value.Value)
	}
	return Result, nil

}

//Modbus 0x04
func (m *ModbusMap) ReadInputRegister(data []byte) ([]byte, error) {
	log.Println("Read Input Register")
	offset := GetOffset(data)
	length := GetLength(data)
	if length > 65535 {
		return nil, fmt.Errorf("out of slice size")
	}
	byteCount := length * 2
	Result := make([]byte, byteCount+1)
	Result[0] = byte(byteCount)
	for i, value := range m.HoldingRegister[offset : length+offset] {
		binary.BigEndian.PutUint16(Result[(i*2)+1:((i+1)*2)+1], value.Value)
	}
	return Result, nil

}

//Modbus (0x05) ForceSingleCoil
func (m *ModbusMap) ForseSingleCoil(data []byte) ([]byte, error) {
	log.Println("Write Single Coil")
	offset := GetOffset(data)
	if len(data) < 2 {
		return nil, fmt.Errorf("not data")
	}
	Result := make([]byte, 4)
	binary.BigEndian.PutUint16(Result[:2], uint16(offset))
	if data[0] == 0xff {
		m.Coil[offset].Value = 1
		Result[2] = 0xff
		Result[3] = 0x00
	} else {
		m.Coil[offset].Value = 0
		Result[2] = 0x00
		Result[3] = 0x00
	}
	return Result, nil
}

//Modbus (0x06) PresetSingleRegister
func (m *ModbusMap) PresetSingleRegister(data []byte) ([]byte, error) {
	log.Println("Write Single Holding Register")
	var value uint16
	length := GetOffset(data)
	value = binary.BigEndian.Uint16(data[2:4])

	m.HoldingRegister[length].Value = value
	newvalue := make([]byte, 4)
	binary.BigEndian.PutUint16(newvalue[:2], uint16(length))
	binary.BigEndian.PutUint16(newvalue[2:], m.HoldingRegister[length].Value)
	return newvalue, nil

}

//Modbus (0x0F)(15) ForseMultipalCoil
/*

 */

func (m *ModbusMap) ForseMultipalCoil(data []byte) ([]byte, error) {
	log.Println("Write MultiCoil")
	offset := GetOffset(data)
	length := GetLength(data)
	bytecount := GetByteCount(data)
	var i, j uint16
	for i = 0; i < bytecount; i++ {
		for j = 0; j < 8; j++ {
			if (((data[5+i]) >> j) % 2) == 1 {
				m.Coil[offset+(i*8)+j].Value = 1
			} else {
				m.Coil[offset+(i*8)+j].Value = 0
			}
		}
	}
	result := make([]byte, 4)
	binary.BigEndian.PutUint16(result[0:2], offset)
	binary.BigEndian.PutUint16(result[2:4], length)
	return result, nil
}

//Modbus (0x10) (16) PresetMultipalRegister
func (m *ModbusMap) PresetMultipalRegister(data []byte) ([]byte, error) {
	log.Println("Write multi Holding Register")
	offset := GetOffset(data)
	length := GetLength(data)
	bytecount := GetByteCount(data)
	var i uint16
	for i = 0; i < length; i++ {
		m.HoldingRegister[offset+i].Value = binary.BigEndian.Uint16(data[5+(i*2) : 5+((i+1)*2)])

	}
	if i*2 != bytecount {
		return nil, fmt.Errorf("write error")
	}
	result := make([]byte, 4)
	binary.BigEndian.PutUint16(result[:2], offset)
	binary.BigEndian.PutUint16(result[2:], length)
	return result, nil
}
