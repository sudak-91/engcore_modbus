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

func GenerateWriteCoilsRequest() []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	data := GenerateCoilMockValues(17)
	unitId := GenerateUnitIDMock()
	bdata, coildatasize := Utility.CoilsConverterSliceValueToUint(data)
	functCode := 15 //Write coils
	_, bDataLength := CalcLengthForMockModbusRawData(coildatasize)
	frame = append(frame, bDataLength...)
	frame = append(frame, unitId, byte(functCode))
	frame = append(frame, bdata...)
	return frame
}

func GenerateReadCoilRequest(RegisterCount int) []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	data := Utility.CreateDataBlockFromReadCoilStatus(0, RegisterCount)
	unitID := GenerateUnitIDMock()
	funcCode := 1
	_, bDatalength := CalcLengthForMockModbusRawData(len(data))
	frame = append(frame, bDatalength...)
	frame = append(frame, unitID, byte(funcCode))
	frame = append(frame, data...)
	return frame
}

func GenerateWriteRegisterRequest() []byte {
	var frame []byte
	_, bresult := GeneratateTransactionIDMock()
	frame = append(frame, bresult...)
	frame = append(frame, 0, 0)
	_, data := GenerateRegistersMock(14, 2568)
	unitId := GenerateUnitIDMock()
	funcCode := 16 //force Multipal registers
	_, bDatalength := CalcLengthForMockModbusRawData(len(data))
	frame = append(frame, bDatalength...)
	frame = append(frame, unitId, byte(funcCode))
	frame = append(frame, data...)
	return frame
}
