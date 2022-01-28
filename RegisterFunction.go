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
			Result[1+k/8] = byte(1 << shift)
		}
	}
	return Result, nil
}

//Modbus 0x02
func readInputStatus(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Read Input Status")
	offset := GetOffset(data)
	length := GetLength(data)
	if offset+length > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register is 65535")
	}
	if length == 0 {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	resultLength := length / 8
	if resultLength%8 != 0 {
		resultLength++
	}
	Result := make([]byte, resultLength+1) // slice for length and result
	Result[0] = byte(resultLength)

	for k, value := range m.inputRegister[offset : offset+length] {
		if value.Value != 0 {
			shift := uint(k) % 8
			Result[1+k/8] = byte(1 << shift)
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
	if offset+length > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	if length == 0 {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	byteCount := length * 2
	Result := make([]byte, byteCount+1)
	Result[0] = byte(byteCount)
	for i, value := range m.holdingRegister[offset : length+offset] {
		binary.BigEndian.PutUint16(Result[i*2:(i+1)*2], value.Value)
	}
	return Result, nil

}

//Modbus 0x04
func readInputRegister(data []byte, m *ModbusRegisters) ([]byte, error) {
	log.Println("Read Input Register")
	offset := GetOffset(data)
	length := GetLength(data)
	if offset+length > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	if length == 0 {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	byteCount := length * 2
	Result := make([]byte, byteCount+1)
	Result[0] = byte(byteCount)
	if offset > length+offset {
		return []byte{ILLEGAL_DATA_VALUE}, fmt.Errorf("illegal data length")
	}
	for i, value := range m.holdingRegister[offset : length+offset] {
		binary.BigEndian.PutUint16(Result[(i*2)+1:((i+1)*2)+1], value.Value)
	}
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
	if data[0] == 0xff {
		m.coil[offset].Value = 1
		Result[2] = 0xff
		Result[3] = 0x00
	} else if data[0] == 0x00 {
		m.coil[offset].Value = 0
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
	var value uint16
	offset := GetOffset(data)
	if offset > 65535 {
		return []byte{ILLEGAL_DATA_ADDRESS}, fmt.Errorf("max register  is 65535")
	}
	value = binary.BigEndian.Uint16(data[2:4])

	m.holdingRegister[offset].Value = value
	newvalue := make([]byte, 4)
	binary.BigEndian.PutUint16(newvalue[:2], uint16(offset))
	binary.BigEndian.PutUint16(newvalue[2:], m.holdingRegister[offset].Value)
	return newvalue, nil

}

//Modbus (0x0F)(15) ForseMultipalCoil
/*

 */

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
			if bitOffset == 1 {
				m.coil[offset+(i*8)+j].Value = 1
			} else if bitOffset == 0 {
				m.coil[offset+(i*8)+j].Value = 0
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
	bytecount := GetByteCount(data)
	var i uint16
	for i = 0; i < length; i++ {
		m.holdingRegister[offset+i].Value = binary.BigEndian.Uint16(data[5+(i*2) : 5+((i+1)*2)])

	}
	if i*2 != bytecount {
		return nil, fmt.Errorf("write error")
	}
	result := make([]byte, 4)
	binary.BigEndian.PutUint16(result[:2], offset)
	binary.BigEndian.PutUint16(result[2:], length)
	return result, nil
}
