package engcore_modbus

const (
	ILLEGAL_FUNCTION      byte = 1
	ILLEGAL_DATA_ADDRESS  byte = 2
	ILLEGAL_DATA_VALUE    byte = 3
	SLAVE_DEVICE_FAILURE  byte = 4
	SERVER_WAIT_SIGNAL    byte = 5
	SERVER_BUSY_SIGNAL    byte = 6
	SERVER_INTERNAL_ERROR byte = 10
)
