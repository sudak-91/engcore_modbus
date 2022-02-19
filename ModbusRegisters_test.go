package engcore_modbus

import (
	"internal/Mock"
	"testing"
)

func TestNewModbusRegisters(t *testing.T) {
	t.Log("----------------------------------------")
	t.Log("Start test \"Create new Modbus Registers\"")
	m, err := NewModbusRegisters()
	if err != nil {
		t.Log("Modbus registers creating fail")
		t.Log(err.Error())
		t.Fail()
	}
	if len(m.coil) != 65535 || len(m.discreteInput) != 65535 ||
		len(m.inputRegister) != 65535 || len(m.holdingRegister) != 65535 {
		t.Log("Register length dont have match")
		t.Fail()
	}
	_, err = NewModbusRegisters(1, 2, 2, 2, 2, 2, 2, 2)
	if err == nil {
		t.Log("Argument over size")
		t.Fail()
	}
	m, err = NewModbusRegisters(10, 10, 10, 10)
	if err != nil {
		t.Log("Modbus registers creating fail")
		t.Log(err.Error())
		t.Fail()
	}
	if len(m.coil) != 10 || len(m.discreteInput) != 10 ||
		len(m.inputRegister) != 10 || len(m.holdingRegister) != 10 {
		t.Log("Register length dont have match")
		t.Fail()
	}
	t.Log("----------------------------------------")

}

func TestGetSetCoilBasic(t *testing.T) {
	t.Log("Start get\\set coils testing")
	m, _ := NewModbusRegisters(10, 10, 10, 10)
	mockresult := Mock.GenerateCoilMockValues(10)
	t.Log("Mock data is:", mockresult)
	err := m.SetCoil(0, mockresult)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range mockresult {
		if int(m.coil[k].Value) != v {
			t.Logf("Not equal %v : %v", m.coil[k], v)
			t.Fail()
		}
	}
	t.Log("Start testing getcoil function")
	getResult, err := m.GetCoil(3, 2)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range getResult {
		if mockresult[3+k] != int(v.Value) {
			t.Logf("Not equal %v : %v", v.Value, mockresult[3+k])
			t.Fail()
		}
	}
}

func TestGetSetDiscreteInputBasic(t *testing.T) {
	t.Log("Start basic Get\\Set Discrete Inputs")
	t.Log("Start basic set test")
	m, _ := NewModbusRegisters(1, 10, 1, 1)
	BasicTestData := Mock.GenerateCoilMockValues(4)
	m.SetDiscreteInput(3, BasicTestData)
	for k, v := range BasicTestData {
		if byte(v) != m.discreteInput[3+k].Value {
			t.Log("Fail data")
			t.Fail()
		}
	}

	t.Log("Start basic get test")
	BasicGetResult, _ := m.GetDiscreteInput(3, 4)
	for k, v := range BasicGetResult {
		if byte(BasicTestData[k]) != v.Value {
			t.Log("Fail data")
			t.Fail()
		}
	}
}

func TestGetSetInputRegistersBasic(t *testing.T) {
	t.Log("Start basic Get\\Set Input Registers")
	m, _ := NewModbusRegisters(1, 1, 15, 1)
	Testdata, _ := Mock.GenerateRegistersMock(12, 3579)
	err := m.SetInputRegister(2, Testdata)
	if err != nil {
		t.Fatal(err.Error())
	}
	for k, v := range Testdata {
		if m.inputRegister[2+k].Value != v {
			t.Logf("Value in register %v is %v != %v ", (2 + k), m.inputRegister[2+k].Value, v)
			t.Fail()
		}
	}
	GetInputRegisrtersData, err := m.GetInputRegister(2, 6)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for k, v := range GetInputRegisrtersData {
		if Testdata[k] != uint16(v.Value) {
			t.Logf("Value in register %v is %v != %v ", (2 + k), v, Testdata[k])
			t.Fail()
		}
	}
}

func TestGetSetHoldingRegistersBasic(t *testing.T) {
	m, err := NewModbusRegisters(1, 1, 1, 23)
	if err != nil {
		t.Fatal(err.Error())
	}
	MockData, _ := Mock.GenerateRegistersMock(12, 5430)
	err = m.SetHoldingRegister(10, MockData)
	if err != nil {
		t.Fatal(err.Error())
	}
	for k, v := range MockData {
		if v != m.holdingRegister[10+k].Value {
			t.Logf("Value in register %v is %v != %v ", (10 + k), m.holdingRegister[10+k].Value, v)
			t.Fail()
		}
	}
	GetHoldingRegisters, _ := m.GetHoldingRegister(10, 12)
	for k, v := range GetHoldingRegisters {
		if v.Value != MockData[k] {
			t.Logf("Value in register %v is %v != %v ", (10 + k), v, MockData[k])
			t.Fail()
		}
	}
}

/*func TestSetCoil(t *testing.T) {
	t.Log("Start get\\set coils testing")
	t.Log("random generation testing cycles count")
	CycleCount := rand.Intn(100)
	t.Logf("Cycle Count is %v", CycleCount)
	time.Sleep(5 * time.Second)
	for i := 0; i < CycleCount; i++ {
		t.Log("Random Generation coil registers length")
		CoilLength := rand.Intn(65535)
		t.Logf("coil length = %v", CoilLength)
		m, err := NewModbusRegisters(CoilLength, 1, 1, 11)
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Log("Generate Mock data size")
		MockDataSize := rand.Intn(CoilLength)
		t.Logf("Mock data size is %v", MockDataSize)
		MockResult := Mock.GenerateCoilMockValues(MockDataSize)
		t.Log("Random Generate Offset")
		CoilOffset := rand.Intn(CoilLength)
		t.Logf("Coil offset is: %v", CoilOffset)
		err = m.SetCoil(CoilOffset, MockResult)
		if err != nil {
			t.Log(err.Error())
			t.Fail()
			continue
		}
		for k, v := range MockResult {
			if int(m.coil[CoilOffset+k].Value) != v {
				t.Logf("Not equal %v :  %v", m.coil[CoilOffset+k].Value, v)
				t.Fail()
			}
		}
	}

}*/
