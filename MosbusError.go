package engcore_modbus

const (
	ILLEGAL_FUNCTION      uint8 = 1
	ILLEGAL_DATA_ADDRESS  uint8 = 2
	ILLEGAL_DATA_VALUE    uint8 = 3
	SLAVE_DEVICE_FAILURE  uint8 = 4
	SERVER_WAIT_SIGNAL    uint8 = 5
	SERVER_BUSY_SIGNAL    uint8 = 6
	SERVER_INTERNAL_ERROR uint8 = 10
)
