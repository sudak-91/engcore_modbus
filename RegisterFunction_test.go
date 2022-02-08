package engcore_modbus

import (
	"bytes"
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
