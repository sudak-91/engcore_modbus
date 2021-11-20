package engcore_modbus

import "fmt"

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
	discreteInput   []ModbusCoil
	coil            []ModbusCoil
	inputRegister   []ModbusRegister
	holdingRegister []ModbusRegister
}

func NewModbusMap() *ModbusMap {
	m := &ModbusMap{}
	m.coil = make([]ModbusCoil, 65535)
	m.discreteInput = make([]ModbusCoil, 65535)
	m.inputRegister = make([]ModbusRegister, 65535)
	m.holdingRegister = make([]ModbusRegister, 65535)
	return m
}

func (m *ModbusMap) GetCoil(offset int) (ModbusCoil, error) {
	if offset > 65535 {
		return ModbusCoil{}, fmt.Errorf("offset oversize")
	}
	return m.coil[offset], nil
}
func (m *ModbusMap) SetCoil(offset int, value byte) error {
	if value > 1 || offset > 65535 {
		return fmt.Errorf("illegal input data")
	}
	m.coil[offset].Value = value
	return nil
}
func (m *ModbusMap) GetDiscreteInput(offset int) (ModbusCoil, error) {
	if offset > 65535 {
		return ModbusCoil{}, fmt.Errorf("offset oversize")
	}
	return m.discreteInput[offset], nil
}
func (m *ModbusMap) SetDiscreteInput(offset int, value byte) error {
	if offset > 65535 || value > 1 {
		return fmt.Errorf("offset oversize")
	}
	m.discreteInput[offset].Value = value
	return nil
}
func (m *ModbusMap) GetInputRegister(offset int) (ModbusRegister, error) {
	if offset > 65535 {
		return ModbusRegister{}, fmt.Errorf("offset oversize")
	}
	return m.inputRegister[offset], nil
}
func (m *ModbusMap) SetInputRegister(offset int, value uint16) error {
	if offset > 65535 {
		return fmt.Errorf("offset oversize")
	}
	m.inputRegister[offset].Value = value
	return nil
}
func (m *ModbusMap) GetHoldingRegister(offset int) (ModbusRegister, error) {
	if offset > 65535 {
		return ModbusRegister{}, fmt.Errorf("offset oversize")
	}
	return m.holdingRegister[offset], nil
}
func (m *ModbusMap) SetHoldingRegister(offset int, value uint16) error {
	if offset > 65535 {
		return fmt.Errorf("offset oversize")
	}
	m.holdingRegister[offset].Value = uint16(value)
	return nil
}
