package engcore_modbus

import (
	"fmt"
	"log"
	"sync"
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

type ModbusRegisters struct {
	discreteMutex        sync.Mutex
	discreteInput        []ModbusCoil
	coilMutex            sync.Mutex
	coil                 []ModbusCoil
	inputRegisterMutex   sync.Mutex
	inputRegister        []ModbusRegister
	holdingRegisterMutex sync.Mutex
	holdingRegister      []ModbusRegister
}

func NewModbusRegisters(RegisterSize ...int) (error, *ModbusRegisters) {
	m := &ModbusRegisters{}
	size := len(RegisterSize)
	RegistersCount := [4]int{65535, 65535, 65535, 65535}
	if size > 4 {
		return fmt.Errorf("register size must have 4 or less elements"), nil
	}
	for k, v := range RegisterSize {
		RegistersCount[k] = v
	}
	m.coil = make([]ModbusCoil, RegistersCount[0])
	m.discreteInput = make([]ModbusCoil, RegistersCount[1])
	m.inputRegister = make([]ModbusRegister, RegistersCount[2])
	m.holdingRegister = make([]ModbusRegister, RegistersCount[3])
	log.Printf("Modbus registers was initializied. Coil's dimensions = %v, Discret input's dimension = %v, Input registers dimension=%v, Holding registers dimension=%v", RegistersCount[0], RegistersCount[1], RegistersCount[2], RegistersCount[3])
	return nil, m
}

//API
func (m *ModbusRegisters) GetCoil(offset int) (ModbusCoil, error) {
	m.coilMutex.Lock()
	defer m.coilMutex.Unlock()
	if offset > len(m.coil) {
		return ModbusCoil{}, fmt.Errorf("coil outside")
	}
	return m.coil[offset], nil
}
func (m *ModbusRegisters) SetCoil(offset int, value byte) error {
	m.coilMutex.Lock()
	defer m.coilMutex.Unlock()
	if value > 1 || offset > len(m.coil) {
		return fmt.Errorf("illegal input data")
	}
	m.coil[offset].Value = value
	return nil
}
func (m *ModbusRegisters) GetDiscreteInput(offset int) (ModbusCoil, error) {
	m.discreteMutex.Lock()
	defer m.discreteMutex.Unlock()
	if offset > len(m.discreteInput) {
		return ModbusCoil{}, fmt.Errorf("discrete input outside")
	}
	return m.discreteInput[offset], nil
}
func (m *ModbusRegisters) SetDiscreteInput(offset int, value byte) error {
	m.discreteMutex.Lock()
	defer m.discreteMutex.Unlock()
	if offset > len(m.discreteInput) || value > 1 {
		return fmt.Errorf("discrete input outside")
	}
	m.discreteInput[offset].Value = value
	return nil
}
func (m *ModbusRegisters) GetInputRegister(offset int) (ModbusRegister, error) {
	m.inputRegisterMutex.Lock()
	defer m.inputRegisterMutex.Unlock()
	if offset > len(m.inputRegister) {
		return ModbusRegister{}, fmt.Errorf("input register outside")
	}
	return m.inputRegister[offset], nil
}
func (m *ModbusRegisters) SetInputRegister(offset int, value uint16) error {
	m.inputRegisterMutex.Lock()
	defer m.inputRegisterMutex.Unlock()
	if offset > len(m.inputRegister) {
		return fmt.Errorf("input register outside")
	}
	m.inputRegister[offset].Value = value
	return nil
}
func (m *ModbusRegisters) GetHoldingRegister(offset int) (ModbusRegister, error) {
	m.holdingRegisterMutex.Lock()
	defer m.holdingRegisterMutex.Unlock()
	if offset > len(m.holdingRegister) {
		return ModbusRegister{}, fmt.Errorf("holding register outside")
	}
	return m.holdingRegister[offset], nil
}
func (m *ModbusRegisters) SetHoldingRegister(offset int, value uint16) error {
	m.holdingRegisterMutex.Lock()
	defer m.holdingRegisterMutex.Unlock()
	if offset > len(m.holdingRegister) {
		return fmt.Errorf("holdin register outside")
	}
	m.holdingRegister[offset].Value = uint16(value)
	return nil
}

func (m *ModbusRegisters) SetCoilValueName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.coil[offset].Description = name
	return nil
}
func (m *ModbusRegisters) SetDiscreteInputName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.discreteInput[offset].Description = name
	return nil
}
func (m *ModbusRegisters) SetInputRegisterName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.inputRegister[offset].Description = name
	return nil
}
func (m *ModbusRegisters) SetHoldingRegisterName(name string, offset int) error {
	if err := ValidateSetNameData(offset, name); err != nil {
		return err
	}
	m.holdingRegister[offset].Description = name
	return nil
}
