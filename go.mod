module github.com/sudak-91/engcore_modbus

go 1.17

require internal/Mock v1.0.0
replace internal/Mock => ./internal/pkg/Mock

require internal/Utility v1.0.0
replace internal/Utility => ./internal/pkg/Utility