package engcore_modbus

import "encoding/binary"

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