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
		return ModbusCoil{}, fmt.Errorf("coil outside")
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
		return ModbusCoil{}, fmt.Errorf("discrete input outside")
	}
	return m.discreteInput[offset], nil
}
func (m *ModbusMap) SetDiscreteInput(offset int, value byte) error {
	if offset > 65535 || value > 1 {
		return fmt.Errorf("discrete input outside")
	}
	m.discreteInput[offset].Value = value
	return nil
}
func (m *ModbusMap) GetInputRegister(offset int) (ModbusRegister, error) {
	if offset > 65535 {
		return ModbusRegister{}, fmt.Errorf("input register outside")
	}
	return m.inputRegister[offset], nil
}
func (m *ModbusMap) SetInputRegister(offset int, value uint16) error {
	if offset > 65535 {
		return fmt.Errorf("input register outside")
	}
	m.inputRegister[offset].Value = value
	return nil
}
func (m *ModbusMap) GetHoldingRegister(offset int) (ModbusRegister, error) {
	if offset > 65535 {
		return ModbusRegister{}, fmt.Errorf("holding register outside")
	}
	return m.holdingRegister[offset], nil
}
func (m *ModbusMap) SetHoldingRegister(offset int, value uint16) error {
	if offset > 65535 {
		return fmt.Errorf("holdin register outside")
	}
	m.holdingRegister[offset].Value = uint16(value)
	return nil
}
func (m *ModbusMap) SetCoilValueName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.coil[offset].Description = name
	return nil
}
func (m *ModbusMap) SetDiscreteInputName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.discreteInput[offset].Description = name
	return nil
}
func (m *ModbusMap) SetInputRegisterName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.inputRegister[offset].Description = name
	return nil
}
func (m *ModbusMap) SetHoldingRegisterName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.holdingRegister[offset].Description = name
	return nil
}
