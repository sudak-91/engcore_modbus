package engcore_modbus

import (
	"testing"
)

func TestRawDataToModbusRawData(t *testing.T) {

	t.Log("Start Test")
	testdata := make([]byte, 512)
	testdata = []byte{00, 01, 00, 00, 00, 06, 01, 02, 00, 00, 00, 02}
	t.Log("Test data is:", testdata)
	rslt := &ModbusRawData{
		TransactionID:  1,
		ProtocolID:     0,
		Length:         6,
		UnitID:         1,
		FunctionalCode: 2,
		Data:           []byte{00, 00, 00, 02},
	}
	testresult, err := RawDataToModbusRawData(testdata)
	if err != nil {
		t.Log("Fatal Error Function")
		t.Fatalf(err.Error())
	}
	if testresult.TransactionID != rslt.TransactionID {
		t.Log("Transcaction corrupted")
		t.Fail()
	}
	if testresult.ProtocolID != rslt.ProtocolID {
		t.Log("Protocol Id fail")
		t.Fail()
	}
	if testresult.Length != rslt.Length {
		t.Log("Length fail")
		t.Fail()
	}
	if testresult.FunctionalCode != rslt.FunctionalCode {
		t.Log("Functional Code fail")
		t.Fail()
	}

	t.Log(testresult.Data)
	t.Log(rslt.Data)

}
