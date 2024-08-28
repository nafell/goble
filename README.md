# Goble
Goble is a Go/C++ based BLE IoT System that monitors room environment values and shows them to users in a graphically interesting manner.
Uses Bluetooth for edge conenction, Socket for backend/frontend data sending.

## How to use
You will need:
- A **PC/server-machine** with Bluetooth capability. Works on Windows/Mac/Linux
- **ESP32 board** (any arduino boards with BLE capability should work)
  - **Arduino IDE** for flashing the program to the board.
- **MHz-19c** CO2 sensor

## Docs
### Architecture (ja_jp)
![architecture diagram](/docs/architecture.drawio.png)
[drawio file](/docs/architecture.drawio)

## Running Sample
Sensor node:  
![Sensor node](/docs/additional_items/mhz19c.webp)

Sensor values(ppm) shown on PC sent via Bluetooth notify signal.  
This will be replaced by a more helpful GUI.  
![sensor log](/docs/additional_items/sensorlog.png)
