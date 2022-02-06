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

func NewModbusRegisters(RegisterSize ...int) (*ModbusRegisters, error) {
	m := &ModbusRegisters{}
	size := len(RegisterSize)
	RegistersCount := [4]int{65535, 65535, 65535, 65535}
	if size > 4 {
		return nil, fmt.Errorf("register size must have 4 or less elements")
	}
	for k, v := range RegisterSize {
		RegistersCount[k] = v
	}
	m.coil = make([]ModbusCoil, RegistersCount[0])
	m.discreteInput = make([]ModbusCoil, RegistersCount[1])
	m.inputRegister = make([]ModbusRegister, RegistersCount[2])
	m.holdingRegister = make([]ModbusRegister, RegistersCount[3])
	log.Printf("Modbus registers was initializied. Coil's dimensions = %v, Discret input's dimension = %v, Input registers dimension=%v, Holding registers dimension=%v", RegistersCount[0], RegistersCount[1], RegistersCount[2], RegistersCount[3])
	return m, nil
}

//API
func (m *ModbusRegisters) GetCoil(offset int, length int) ([]ModbusCoil, error) {
	m.coilMutex.Lock()
	defer m.coilMutex.Unlock()
	if length == 0 {
		return nil, fmt.Errorf("zero length")
	}
	if offset+length > len(m.coil) {
		return nil, fmt.Errorf("coil outside")
	}
	return m.coil[offset : offset+length], nil
}

func (m *ModbusRegisters) SetCoil(offset int, value []int) error {
	m.coilMutex.Lock()
	defer m.coilMutex.Unlock()
	if offset+len(value) > len(m.coil) {
		return fmt.Errorf("length over size")
	}
	for k, v := range value {
		if v > 1 {
			return fmt.Errorf("illegal data")
		}
		m.coil[offset+k].Value = byte(v)

	}
	return nil
}

func (m *ModbusRegisters) GetDiscreteInput(offset int, length int) ([]ModbusCoil, error) {
	m.discreteMutex.Lock()
	defer m.discreteMutex.Unlock()
	if offset+length > len(m.discreteInput) {
		return nil, fmt.Errorf("discrete input outside")
	}
	return m.discreteInput[offset : offset+length], nil
}

func (m *ModbusRegisters) SetDiscreteInput(offset int, value []int) error {
	m.discreteMutex.Lock()
	defer m.discreteMutex.Unlock()
	if offset+len(value) > len(m.discreteInput) {
		return fmt.Errorf("discrete input outside")
	}
	for k, v := range value {
		if v > 1 {
			return fmt.Errorf("illegal data")
		}
		m.discreteInput[offset+k].Value = byte(v)
	}

	return nil
}

func (m *ModbusRegisters) GetInputRegister(offset int, length int) ([]ModbusRegister, error) {
	m.inputRegisterMutex.Lock()
	defer m.inputRegisterMutex.Unlock()
	if offset+length > len(m.inputRegister) {
		return nil, fmt.Errorf("input register outside")
	}
	return m.inputRegister[offset : offset+length], nil
}

func (m *ModbusRegisters) SetInputRegister(offset int, value []uint16) error {
	m.inputRegisterMutex.Lock()
	defer m.inputRegisterMutex.Unlock()
	if offset+len(value) > len(m.inputRegister) {
		return fmt.Errorf("input register outside")
	}
	for k, v := range value {
		m.inputRegister[offset+k].Value = v
	}
	return nil
}

func (m *ModbusRegisters) GetHoldingRegister(offset int, length int) ([]ModbusRegister, error) {
	m.holdingRegisterMutex.Lock()
	defer m.holdingRegisterMutex.Unlock()
	if offset+length > len(m.holdingRegister) {
		return nil, fmt.Errorf("holding register outside")
	}
	return m.holdingRegister[offset : offset+length], nil
}

func (m *ModbusRegisters) SetHoldingRegister(offset int, value []uint16) error {
	m.holdingRegisterMutex.Lock()
	defer m.holdingRegisterMutex.Unlock()
	if offset+len(value) > len(m.holdingRegister) {
		return fmt.Errorf("holdin register outside")
	}
	for k, v := range value {
		m.holdingRegister[offset+k].Value = v
	}
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

func (m *ModbusRegisters) GetCoilLength() int {
	return len(m.coil)
}

func (m *ModbusRegisters) GetDiscreteInputLength() int {
	return len(m.discreteInput)
}
