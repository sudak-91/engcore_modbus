package Utility

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
		var offset int
		offset = k
		if k > 8 {
			offset = k - 8
		}
		if v == 1 {
			result[k/8] = result[k/8] | byte(1<<offset)
		}
	}
	return
}
