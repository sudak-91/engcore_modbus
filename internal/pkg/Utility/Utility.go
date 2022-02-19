package Utility

import "encoding/binary"

func OneCoilConvertIntToByteSlice(Value int) []byte {
	ByteValue := []byte{0x00, 0x00}
	if Value == 1 {
		ByteValue[0] = 0xFF
	}
	return ByteValue
}

func CoilsConverterSliceValueToUint(Value []int) (result []byte, size int) {
	size = len(Value) / 8
	if len(Value)%8 != 0 {
		size++
	}
	result = make([]byte, size)
	for k, v := range Value {

		offset := k % 8

		if v == 1 {
			result[k/8] = result[k/8] | byte(1<<offset)
		}
	}
	return
}

func ConvertOffsetFromIntToByteSlice(offset int) []byte {
	bResult := make([]byte, 2)
	binary.BigEndian.PutUint16(bResult, uint16(offset))
	return bResult
}

func ConvertLengthfromIntToByteSlice(length int) []byte {
	bResult := make([]byte, 2)
	binary.BigEndian.PutUint16(bResult, uint16(length))
	return bResult
}

func CreateDataByteSliceForModbusDataFrameToReadFunction(offset, length int) []byte {
	var bResult []byte
	bResult = append(bResult, ConvertOffsetFromIntToByteSlice(offset)...)
	bResult = append(bResult, ConvertLengthfromIntToByteSlice(length)...)
	return bResult
}

func ConvertUintSliceToByteSlice(source []uint16) []byte {
	var result []byte
	temp := make([]byte, 2)
	for _, v := range source {

		binary.BigEndian.PutUint16(temp[:2], v)
		result = append(result, temp...)
	}
	return result
}
