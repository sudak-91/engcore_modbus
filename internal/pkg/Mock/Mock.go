package Mock

import (
	"encoding/binary"
	"internal/Utility"
	"math/rand"
)

func GenerateCoilMockValues(coilscount int) []int {
	resultslice := make([]int, coilscount)

	for k, _ := range resultslice {
		resultslice[k] = rand.Intn(2)
	}
	return resultslice

}

func GenerateRegistersMock(registercount int, maxvalue int) ([]uint16, []byte) {
	resultslice := make([]uint16, registercount)
	btemp := make([]byte, 2)
	var bResultslice []byte
	for k, _ := range resultslice {
		resultslice[k] = uint16(rand.Intn(maxvalue))
		binary.BigEndian.PutUint16(btemp, resultslice[k])
		bResultslice = append(bResultslice, btemp...)
	}
	return resultslice, bResultslice
}

func GenerateRegistersOffset(MaxRegisterSize int) (OffsetByte []byte, OffsetInt int) {
	OffsetAddressInt := rand.Intn(MaxRegisterSize)
	OffsetInt = OffsetAddressInt
	OffsetByte = make([]byte, 2)
	binary.BigEndian.PutUint16(OffsetByte, uint16(OffsetAddressInt))
	return
}

func GeneratateTransactionIDMock() (iTransactionID uint16, bTransactionID []byte) {
	iTransactionID = uint16(rand.Intn(10))
	bTransactionID = make([]byte, 2)
	binary.BigEndian.PutUint16(bTransactionID, iTransactionID)
	return
}
func CalcLengthForMockModbusRawData(datalen int) (iDatalength uint16, bDataLength []byte) {
	iDatalength = uint16(2 + datalen)
	bDataLength = make([]byte, 2)
	binary.BigEndian.PutUint16(bDataLength, iDatalength)
	return
}

func GenerateUnitIDMock() (UnitID byte) {
	UnitID = byte(rand.Intn(256))
	return
}

/*Generate ModbusTCP command for write manu coils
 */
func GenerateWriteCoilsRequest(offset int, length int) []byte {
	var frame []byte
	functCode := 0x0F //Write coils

	_, bTransactionID := GeneratateTransactionIDMock()
	frame = append(frame, bTransactionID...)
	//Add to frame protocol identify 0x00 0x00
	frame = append(frame, 0, 0)

	bOffset := make([]byte, 2)
	binary.BigEndian.PutUint16(bOffset, uint16(offset))

	bRegistersCount := make([]byte, 2)
	binary.BigEndian.PutUint16(bRegistersCount, uint16(length))

	data := GenerateCoilMockValues(length)
	bdata, coildatasize := Utility.CoilsConverterSliceValueToUint(data)

	unitId := GenerateUnitIDMock()

	_, bDataLength := CalcLengthForMockModbusRawData(coildatasize + 5)

	frame = append(frame, bDataLength...)
	frame = append(frame, unitId, byte(functCode))

	frame = append(frame, bOffset...)
	frame = append(frame, bRegistersCount...)
	frame = append(frame, byte(coildatasize))
	frame = append(frame, bdata...)
	return frame
}

func GenerateWriteSingleCoil(offset int, value int) []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	data := Utility.OneCoilConvertIntToByteSlice(value)
	unitId := GenerateUnitIDMock()
	funcCode := 5
	_, bDataLength := CalcLengthForMockModbusRawData(len(data) + 2)
	frame = append(frame, bDataLength...)
	frame = append(frame, unitId, byte(funcCode))
	boffset := make([]byte, 2)
	binary.BigEndian.PutUint16(boffset, uint16(offset))
	frame = append(frame, boffset...)
	frame = append(frame, data...)
	return frame
}

func GenerateWriteSingleRegister(offset int, maxvalue int) []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	_, bDataRegister := GenerateRegistersMock(1, maxvalue)
	unitId := GenerateUnitIDMock()
	funcCode := 6
	_, bDataLength := CalcLengthForMockModbusRawData(len(bDataRegister) + 2)
	frame = append(frame, bDataLength...)
	frame = append(frame, unitId, byte(funcCode))
	boffset := make([]byte, 2)
	binary.BigEndian.PutUint16(boffset, uint16(offset))
	frame = append(frame, boffset...)
	frame = append(frame, bDataRegister...)
	return frame
}

func GenerateReadCoilRequest(RegisterCount, functionalCode int) []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	data := Utility.CreateDataByteSliceForModbusDataFrameToReadFunction(0, RegisterCount)
	unitID := GenerateUnitIDMock()
	funcCode := functionalCode
	_, bDatalength := CalcLengthForMockModbusRawData(len(data))
	frame = append(frame, bDatalength...)
	frame = append(frame, unitID, byte(funcCode))
	frame = append(frame, data...)
	return frame
}

func GenerateReadRegisters(RegisterCount, functionalCode int) []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	data := Utility.CreateDataByteSliceForModbusDataFrameToReadFunction(0, RegisterCount)
	unitID := GenerateUnitIDMock()
	_, bDataLength := CalcLengthForMockModbusRawData(len(data))
	frame = append(frame, bDataLength...)
	frame = append(frame, unitID, byte(functionalCode))
	frame = append(frame, data...)
	return frame

}

func GenerateWriteRegisterRequest(offset int, registersCount int) []byte {
	var frame []byte
	_, bTransactionID := GeneratateTransactionIDMock()
	frame = append(frame, bTransactionID...)
	frame = append(frame, 0, 0)

	_, data := GenerateRegistersMock(registersCount, 2568)
	unitId := GenerateUnitIDMock()
	funcCode := 16 //force Multipal registers

	bOffset := make([]byte, 2)
	binary.BigEndian.PutUint16(bOffset, uint16(offset))
	bRegisterCount := make([]byte, 2)
	binary.BigEndian.PutUint16(bRegisterCount, uint16(registersCount))

	_, bDatalength := CalcLengthForMockModbusRawData(len(data) + 5)
	frame = append(frame, bDatalength...)
	frame = append(frame, unitId, byte(funcCode))
	frame = append(frame, bOffset...)
	frame = append(frame, bRegisterCount...)
	frame = append(frame, byte(len(data)))
	frame = append(frame, data...)
	return frame
}
