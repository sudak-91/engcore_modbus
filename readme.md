# ModbusTCP Server
ModbusTCP сервер и API для работы с регистрами.

# Команды поддерживаемые ModbusTCP Server
1. 0x01 — чтение значений из нескольких регистров флагов (Read Coil Status).
2. 0x02 — чтение значений из нескольких дискретных входов (Read Discrete Inputs).
3. 0x03 — чтение значений из нескольких регистров хранения (Read Holding Registers).
4. 0x04 — чтение значений из нескольких регистров ввода (Read Input Registers).
5. 0x05 — запись значения одного флага (Force Single Coil).
6. 0x06 — запись значения в один регистр хранения (Preset Single Register).
1. 0x0F — запись значений в несколько регистров флагов (Force Multiple Coils)
1. 0x10 — запись значений в несколько регистров хранения (Preset Multiple Registers)

**Все регистры имеют формат uint16**
