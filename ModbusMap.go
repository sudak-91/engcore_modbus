package engcore_modbus

import (
	"encoding/binary"
	"fmt"
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
}

func NewModbusMap() *ModbusMap {
	m := &ModbusMap{}
	m.Coil = make([]ModbusCoil, 65535)
	m.DiscreteInput = make([]ModbusCoil, 65535)
	m.InputRegister = make([]ModbusRegister, 65535)
	m.HoldingRegister = make([]ModbusRegister, 65535)
	return m
}

//Return: data, error
//Mosbuc Command 0x01
func (m *ModbusMap) ReadCoilStatus(offset int, RegisterCount int) ([]byte, error) {
	if RegisterCount > 65535 {
		return nil, fmt.Errorf("max coil count is 65535")
	}

	resultLength := RegisterCount / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1) // slice for length and result
	Result[0] = byte(resultLength)

	data := m.Coil[offset : offset+RegisterCount]
	for k, value := range data {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = byte(1 << shift)
		}

	}
	return Result, nil

}

//Modbus 0x02
func (m *ModbusMap) ReadInputStatus(offset int, length int) ([]byte, error) {
	if length > 65535 {
		return nil, fmt.Errorf("max coil count is 65535")
	}

	resultLength := length / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1) // slice for length and result
	Result[0] = byte(resultLength)

	data := m.InputRegister[offset : offset+length]
	for k, value := range data {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = byte(1 << shift)
		}

	}
	return Result, nil
}

//Modbus 0x03
func (m *ModbusMap) ReadHoldingRegisters(offset int, length int) ([]byte, error) {
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
func (m *ModbusMap) ReadInputRegister(offset int, length int) ([]byte, error) {
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

//Modbus (0x05) ForceSingleCoil
func (m *ModbusMap) ForseSingleCoil(offset int, inputdata []byte) ([]byte, error) {
	if len(inputdata) > 2 {
		return nil, fmt.Errorf("Not data")
	}
	Result := make([]byte, 4)
	binary.BigEndian.PutUint16(Result[:2], uint16(offset))
	if inputdata[0] == 0xff {
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

//Modbus (0x0F)(15) ForseMultipalCoil

//Modbus (0x10) (16) PresetMultipalRegister
