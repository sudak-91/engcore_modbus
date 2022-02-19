package engcore_modbus

import (
	"encoding/binary"
	"fmt"
	"log"
)

//Modbuc Command 0x01
func readCoilStatus(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Read Coil")
	offset := GetOffset(data)
	length := GetLength(data)
	CoilResult, err := m.GetCoil(int(offset), int(length)) // slice for length and result
	if err != nil {

		return []byte{ILLEGAL_DATA_ADDRESS}, err
	}
	resultLength := length / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1)
	Result[0] = byte(resultLength)

	for k, value := range CoilResult {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = Result[1+k/8] | byte(1<<shift)
		}
	}
	return Result, nil
}

//Modbus 0x02
func readInputStatus(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Read Input Status")
	offset := GetOffset(data)
	length := GetLength(data)
	InputResult, err := m.GetDiscreteInput(int(offset), int(length))
	if err != nil {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}

	resultLength := length / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1) // slice for length and result
	Result[0] = byte(resultLength)

	for k, value := range InputResult {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = Result[1+k/8] | byte(1<<shift)
		}

	}
	return Result, nil
}

//HoldingRegister
//Modbus 0x03
func readHoldingRegisters(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Read Holding Register")
	offset := GetOffset(data)
	length := GetLength(data)
	HoldingRegistersResult, err := m.GetHoldingRegister(int(offset), int(length))

	if err != nil {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}

	byteCount := length * 2
	Result := make([]byte, byteCount)
	for i, value := range HoldingRegistersResult {
		binary.BigEndian.PutUint16(Result[i*2:(i*2)+2], value.Value)
	}
	b := make([]byte, 1)
	b[0] = byte(byteCount)
	Result = append(b, Result...)
	return Result, nil

}

//Modbus 0x04
func readInputRegister(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Read Input Register")
	offset := GetOffset(data)
	length := GetLength(data)
	InputRegistersResult, err := m.GetInputRegister(int(offset), int(length))
	if err != nil {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	byteCount := length * 2
	Result := make([]byte, byteCount)
	for i, value := range InputRegistersResult {
		binary.BigEndian.PutUint16(Result[i*2:(i*2)+2], value.Value)
	}
	b := make([]byte, 1)
	b[0] = byte(byteCount)
	Result = append(b, Result...)
	return Result, nil

}

//Modbus (0x05) ForceSingleCoil
func forseSingleCoil(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Write Single Coil")
	offset := GetOffset(data)
	if offset > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	if len(data) < 2 {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	Result := make([]byte, 4)
	binary.BigEndian.PutUint16(Result[:2], uint16(offset))
	if data[2] == 0xff {
		m.SetCoil(int(offset), []int{1})
		Result[2] = 0xff
		Result[3] = 0x00
	} else if data[2] == 0x00 {
		m.SetCoil(int(offset), []int{0})
		Result[2] = 0x00
		Result[3] = 0x00
	} else {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal coil data")
	}
	return Result, nil
}

//Modbus (0x06) presetSingleRegister
func presetSingleRegister(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Write Single Holding Register")
	offset := GetOffset(data)
	if offset > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	value := make([]uint16, 1)
	value[0] = binary.BigEndian.Uint16(data[2:4])
	m.SetHoldingRegister(int(offset), value)
	variable, _ := m.GetHoldingRegister(int(offset), 1)
	newvalue := make([]byte, 4)
	binary.BigEndian.PutUint16(newvalue[:2], uint16(offset))
	binary.BigEndian.PutUint16(newvalue[2:], variable[0].Value)
	return newvalue, nil

}

//Modbus (0x0F)(15) ForseMultipalCoil

func forseMultipalCoil(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Write MultiCoil")
	offset := GetOffset(data)
	length := GetLength(data)
	if offset+length > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	if length == 0 {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	bytecount := GetByteCount(data)
	if bytecount == 0 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("byte count equl 0")
	}
	var i, j uint16
	for i = 0; i < bytecount; i++ {
		for j = 0; j < 8; j++ {
			bitOffset := ((data[5+i]) >> j) % 2
			CoilOffset := offset + (i * 8) + j
			if bitOffset == 1 {
				m.SetCoil(int(CoilOffset), []int{1})
			}
		}
	}
	result := make([]byte, 4)
	binary.BigEndian.PutUint16(result[0:2], offset)
	binary.BigEndian.PutUint16(result[2:4], length)
	return result, nil
}

//Modbus (0x10) (16) presetMultipalRegister
func presetMultipalRegister(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Write multi Holding Register")
	offset := GetOffset(data)
	length := GetLength(data)
	if offset+length > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	if length == 0 {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}

	convertresult := ByteSliceToUintSlise(data[5:])
	log.Printf("covert result is: %v", convertresult)
	m.SetHoldingRegister(int(offset), convertresult)
	result := make([]byte, 4)
	binary.BigEndian.PutUint16(result[:2], offset)
	binary.BigEndian.PutUint16(result[2:], length)
	return result, nil
}
