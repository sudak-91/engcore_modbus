package engcore_modbus

import (
	"encoding/binary"
	"fmt"
)

//First register offset
func GetOffset(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[:2])
}

//Count of register
func GetLength(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[2:4])
}

//Count of data bytes
func GetByteCount(data []byte) uint16 {
	result := make([]byte, 2)
	result[0] = 0
	result[1] = data[4]
	return binary.BigEndian.Uint16(result)
}

//Check offset and name length
func ValidateSetNameData(offset int, name string) error {
	if offset > 65535 {
		return fmt.Errorf("coil outside")
	}
	if len(name) == 0 {
		return fmt.Errorf("empty name")
	}
	return nil
}
