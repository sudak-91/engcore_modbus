package engcore_modbus

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
