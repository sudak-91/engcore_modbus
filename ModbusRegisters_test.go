package engcore_modbus

import (
	"internal/Mock"
	"testing"
)

func TestNewModbusRegisters(t *testing.T) {
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

}

func TestGetSetCoil(t *testing.T) {
	t.Log("Start get\\set cils testing")
	m, _ := NewModbusRegisters(10, 10, 10, 10)
	mockresult := Mock.GenerateCoilMoсk(10)
	t.Log("Mock data is:", mockresult)
	err := m.SetCoil(0, []byte{1})
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range mockresult {
		if m.coil[k].Value != v {
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
		if mockresult[3+k] != v.Value {
			t.Logf("Not equal %v : %v", v.Value, mockresult[3+k])
			t.Fail()
		}
	}
}
