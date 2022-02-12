package engcore_modbus

import (
	"bytes"
	"encoding/binary"
	"internal/Mock"
	"internal/Utility"
	"testing"
)

func TestReadCoilStatus(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(20, 20, 20, 20)
	Coilsdata := Mock.GenerateCoilMockValues(10)
	t.Logf("coils data was generate: %v\n", Coilsdata)
	_ = ModbusRegisters.SetCoil(0, Coilsdata)
	rsl, _ := ModbusRegisters.GetCoil(0, 10)
	for _, v := range rsl {
		t.Log(v.Value)
	}
	//@TODO:
	ModbusRequest := Mock.GenerateReadCoilRequest(10, 1)
	t.Logf("ReadCoilRequest was generate: %v\n", ModbusRequest)
	ModbusFrame, _ := RawDataToModbusRawData(ModbusRequest)
	t.Logf("Modbus frame is %v", ModbusFrame)
	t.Logf("Modbus frame Data is %v\n", ModbusFrame.Data)

	TestfunctionResult, err := readCoilStatus(ModbusFrame.Data, ModbusRegisters)
	t.Logf("Read coil status result is %v\n", TestfunctionResult)
	if err != nil {
		t.Fatal()
	}
	byteCoilsData, _ := Utility.CoilsConverterSliceValueToUint(Coilsdata)
	t.Logf("Coil data to byte slice is %v\n", byteCoilsData)
	if !bytes.Equal(TestfunctionResult[1:], byteCoilsData) {
		t.Fatal("Not eqal")
	}
	t.Log("is equal")

}

func TestReadInputStatus(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(20, 20, 20, 20)
	CoilsData := Mock.GenerateCoilMockValues(10)
	t.Logf("coils data was generate: %v\n", CoilsData)
	_ = ModbusRegisters.SetDiscreteInput(0, CoilsData)

	ModbusRequest := Mock.GenerateReadCoilRequest(10, 2)
	t.Logf("ReadCoilRequest was generate: %v\n", ModbusRequest)
	ModbusFrame, _ := RawDataToModbusRawData(ModbusRequest)
	t.Logf("Modbus frame is %v", ModbusFrame)
	t.Logf("Modbus frame Data is %v\n", ModbusFrame.Data)
	TestfunctionResult, err := readInputStatus(ModbusFrame.Data, ModbusRegisters)
	t.Logf("Read coil status result is %v\n", TestfunctionResult)
	if err != nil {
		t.Fatal()
	}
	byteCoilsData, _ := Utility.CoilsConverterSliceValueToUint(CoilsData)
	t.Logf("Coil data to byte slice is %v\n", byteCoilsData)
	if !bytes.Equal(TestfunctionResult[1:], byteCoilsData) {
		t.Fatal("Not eqal")
	}
	t.Log("is equal")

}
func TestReadHoldingRegisters(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(20, 20, 20, 20)
	RegistersData, _ := Mock.GenerateRegistersMock(10, 13456)
	t.Logf("coils data was generate: %v\n", RegistersData)
	_ = ModbusRegisters.SetHoldingRegister(0, RegistersData)
	ModbusRequest := Mock.GenerateReadRegisters(10, 3)
	ModbusFrame, _ := RawDataToModbusRawData(ModbusRequest)
	t.Logf("Modbus frame is %v", ModbusFrame)
	t.Logf("Modbus frame Data is %v\n", ModbusFrame.Data)
	TestfunctionalResult, err := readHoldingRegisters(ModbusFrame.Data, ModbusRegisters)
	t.Logf("Read coil status result is %v\n", TestfunctionalResult)
	if err != nil {
		t.Fatal()
	}
	bData := Utility.ConvertUintSliceToByteSlice(RegistersData)
	t.Logf("Registers data to byte slice is %v\n", bData)
	t.Logf("Registers data to byte slice is %v\n", TestfunctionalResult[1:])
	if !bytes.Equal(TestfunctionalResult[1:], bData) {
		t.Fatal("Not eqal")
	}
	t.Log("is equal")
}

func TestReadInputRegisters(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(20, 20, 20, 20)
	RegistersData, _ := Mock.GenerateRegistersMock(10, 13456)
	t.Logf("coils data was generate: %v\n", RegistersData)
	_ = ModbusRegisters.SetInputRegister(0, RegistersData)
	ModbusRequest := Mock.GenerateReadRegisters(10, 4)
	ModbusFrame, _ := RawDataToModbusRawData(ModbusRequest)
	t.Logf("Modbus frame is %v", ModbusFrame)
	t.Logf("Modbus frame Data is %v\n", ModbusFrame.Data)
	TestfunctionalResult, err := readInputRegister(ModbusFrame.Data, ModbusRegisters)
	t.Logf("Read coil status result is %v\n", TestfunctionalResult)
	if err != nil {
		t.Fatal()
	}
	bData := Utility.ConvertUintSliceToByteSlice(RegistersData)
	t.Logf("Registers data to byte slice is %v\n", bData)
	t.Logf("Registers data to byte slice is %v\n", TestfunctionalResult[1:])
	if !bytes.Equal(TestfunctionalResult[1:], bData) {
		t.Fatal("Not eqal")
	}
	t.Log("is equal")
}

func TestForceSingleCoil(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(10, 10, 10, 10)
	ModbusBytes := Mock.GenerateWriteSingleCoil(2, 1)
	t.Logf("Generator post data: %v\n", ModbusBytes)
	ModbusFrame, _ := RawDataToModbusRawData(ModbusBytes)
	t.Logf("ModbusData is: %v\n", ModbusFrame.Data[2])
	Testrslt, err := forseSingleCoil(ModbusFrame.Data, ModbusRegisters)
	if err != nil {
		t.Fatal()
	}
	t.Logf("result force single coil is %v", Testrslt)
	rslt, _ := ModbusRegisters.GetCoil(2, 1)
	byterslt := Utility.OneCoilConvertIntToByteSlice(int(rslt[0].Value))
	if rslt[0].Value == 0 {
		t.Fail()
		return
	}
	if !bytes.Equal(byterslt, Testrslt[2:]) {
		t.Fail()
		return
	}

}

func TestPresetSingleRegister(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(10, 10, 10, 10)
	ModbusBytes := Mock.GenerateWriteSingleRegister(0, 3424)
	t.Logf("Generator post data: %v\n", ModbusBytes)
	ModbusFrame, _ := RawDataToModbusRawData(ModbusBytes)
	t.Logf("ModbusData is: %v\n", ModbusFrame.Data[2])
	Testrslt, err := presetSingleRegister(ModbusFrame.Data, ModbusRegisters)
	if err != nil {
		t.Fatal()
	}
	t.Logf("result force single coil is %v", Testrslt)
	rslt, _ := ModbusRegisters.GetHoldingRegister(0, 1)
	byteresult := make([]byte, 2)
	binary.BigEndian.PutUint16(byteresult, rslt[0].Value)
	if !bytes.Equal(byteresult, Testrslt[2:]) {
		t.Fail()
		return
	}

}

func TestForceMultipalCoil(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(10, 20, 10, 10)
	ModbusWriteCoilsBytes := Mock.GenerateWriteCoilsRequest(0, 10)
	t.Logf("Generator post data: %v\n", ModbusWriteCoilsBytes)
	ModbusFrame, err := RawDataToModbusRawData(ModbusWriteCoilsBytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ModbusData is: %v\n", ModbusFrame)
	t.Logf("ModbuData \"Data\" field is: %v\n", ModbusFrame.Data)

	Testrslt, err := forseMultipalCoil(ModbusFrame.Data, ModbusRegisters)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Force Multipal Coil Result: %v\n", Testrslt)
	rslt, err := ModbusRegisters.GetCoil(0, 10)
	if err != nil {
		t.Fatal(err)
	}
	var irslt []int
	for _, v := range rslt {
		irslt = append(irslt, int(v.Value))
	}
	brslt, _ := Utility.CoilsConverterSliceValueToUint(irslt)
	t.Logf("Modbus result: %v\n", ModbusFrame.Data)
	t.Logf("Bytes result: %v\n", brslt)
	ModbusReadCoilsBytes := Mock.GenerateReadCoilRequest(10, 2)
	t.Logf("MOdbus Read bytes is: %v", ModbusReadCoilsBytes)
	ModbusReadFrame, _ := RawDataToModbusRawData(ModbusReadCoilsBytes)
	ReadResult, _ := readCoilStatus(ModbusReadFrame.Data, ModbusRegisters)
	t.Logf("ReadResult is: %v", ReadResult[1:])
	if !bytes.Equal(brslt, ReadResult[1:]) {
		t.Fail()
		return
	}

}

func TestPresetMultipalRegister(t *testing.T) {
	ModbusRegisters, _ := NewModbusRegisters(20, 20, 20, 20)
	ModbusWriteRegisterButes := Mock.GenerateWriteRegisterRequest(0, 10)
	t.Logf("Generator post data: %v\n", ModbusWriteRegisterButes)
	ModbusWriteFrame, err := RawDataToModbusRawData(ModbusWriteRegisterButes)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("Modbus Write frame: %v\n", ModbusWriteFrame)
	testResult, err := presetMultipalRegister(ModbusWriteFrame.Data, ModbusRegisters)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Test result is: %v\n", testResult)

	ModbusReadRegisterButes := Mock.GenerateReadRegisters(10, 3)
	t.Logf("Modbus read registers bytes: %v", ModbusReadRegisterButes)
	ModbusReadFrame, err := RawDataToModbusRawData(ModbusReadRegisterButes)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("Modbus Read Frame: %v\n", ModbusReadFrame)
	readtestResult, _ := readHoldingRegisters(ModbusReadFrame.Data, ModbusRegisters)
	t.Logf("Read test result: %v", readtestResult)

	strResult, _ := ModbusRegisters.GetHoldingRegister(0, 10)
	var bResult []byte
	buffer := make([]byte, 2)
	for _, v := range strResult {
		binary.BigEndian.PutUint16(buffer, v.Value)
		bResult = append(bResult, buffer...)
	}
	t.Logf("First object: %v\n", readtestResult[1:])
	t.Logf("Second object: %v\n", bResult)
	if !bytes.Equal(readtestResult[1:], bResult) {
		t.Fail()
		return
	}
}
