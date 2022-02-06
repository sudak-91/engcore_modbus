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
	ModbusRequest := Mock.GenerateReadCoilRequest(10)
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
