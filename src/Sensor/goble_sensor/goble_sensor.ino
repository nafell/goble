/*
    Video: https://www.youtube.com/watch?v=oCMOYS71NIU
    Based on Neil Kolban example for IDF: https://github.com/nkolban/esp32-snippets/blob/master/cpp_utils/tests/BLE%20Tests/SampleNotify.cpp
    Ported to Arduino ESP32 by Evandro Copercini

   Create a BLE server that, once we receive a connection, will send periodic notifications.
   The service advertises itself as: 6E400001-B5A3-F393-E0A9-E50E24DCCA9E
   Has a characteristic of: 6E400002-B5A3-F393-E0A9-E50E24DCCA9E - used for receiving data with "WRITE"
   Has a characteristic of: 6E400003-B5A3-F393-E0A9-E50E24DCCA9E - used to send data with  "NOTIFY"

   The design of creating the BLE server is:
   1. Create a BLE Server
   2. Create a BLE Service
   3. Create a BLE Characteristic on the Service
   4. Create a BLE Descriptor on the characteristic
   5. Start the service.
   6. Start advertising.

   In this example rxValue is the data received (only accessible inside that function).
   And txValue is the data to be sent, in this example just a byte incremented every second.
*/
#include <BLEDevice.h>
#include <BLEServer.h>
#include <BLEUtils.h>
#include <BLE2902.h>

#include "Mhz19_uart.h"
#include "Mhz19_bytes.h"
#include <SoftwareSerial.h>

#define ARDUINO_ARCH_ESP32

#define RXPIN D3
#define TXPIN D2
#define MHZ19C_BAUD 9600
Mhz19c mhz19c;
EspSoftwareSerial::UART uartPort;

uint8_t ppm[2] = {0xFF, 0xFF};

BLEServer *pServer = NULL;
BLECharacteristic * pTxCharacteristic;
bool deviceConnected = false;
bool oldDeviceConnected = false;
uint8_t txValue = 0;

// See the following for generating UUIDs:
// https://www.uuidgenerator.net/

#define SERVICE_UUID           "E3200001-c577-4615-bc4a-44feb3e806fd" // UART service UUID
#define CHARACTERISTIC_UUID_RX "E3200002-c577-4615-bc4a-44feb3e806fd"
#define CHARACTERISTIC_UUID_TX "E3200003-c577-4615-bc4a-44feb3e806fd"


class MyServerCallbacks: public BLEServerCallbacks {
    void onConnect(BLEServer* pServer) {
      deviceConnected = true;
    };

    void onDisconnect(BLEServer* pServer) {
      deviceConnected = false;
    }
};

class MyCallbacks: public BLECharacteristicCallbacks {
    void onWrite(BLECharacteristic *pCharacteristic) {
      std::string rxValue = pCharacteristic->getValue();

      if (rxValue.length() > 0) {
        Serial.println("*********");
        Serial.print("Received Value: ");
        for (int i = 0; i < rxValue.length(); i++)
          Serial.print(rxValue[i]);

        Serial.println();
        Serial.println("*********");
      }
    }
};

void led_init() {
  pinMode(LED_RED, OUTPUT);
  pinMode(LED_GREEN, OUTPUT);
  pinMode(LED_BLUE, OUTPUT);
  pinMode(LED_BUILTIN, OUTPUT);
}

void led_rgb(uint8_t r, uint8_t g, uint8_t b) {
  digitalWrite(LED_BUILTIN, HIGH);
  digitalWrite(LED_RED, r);
  digitalWrite(LED_GREEN, g);
  digitalWrite(LED_BLUE, b);
}

void led_off() {
  digitalWrite(LED_BUILTIN, HIGH);
  digitalWrite(LED_RED, HIGH);
  digitalWrite(LED_GREEN, HIGH);
  digitalWrite(LED_BLUE, HIGH);
}


void setup() {

  Serial.begin(115200);

  // initialize LED pins
  led_init();
  // LED indicate initialization
  led_rgb(LOW, HIGH, HIGH);

  // sensor setup
  uartPort.begin(MHZ19C_BAUD, SWSERIAL_8N1, RXPIN, TXPIN, false);
  if (!uartPort) { // If the object did not initialize, then its configuration is invalid
    Serial.println("Invalid EspSoftwareSerial pin configuration, check config");
    while (1) { // Don't continue with invalid configuration
      delay(1000);
    }
  }

  led_rgb(LOW, LOW, HIGH);

  mhz19c.begin(&uartPort);
  mhz19c.setMeasuringRange(Mhz19MeasuringRange::Ppm_5000);
  mhz19c.enableAutoBaseCalibration();

  Serial.println("Preheating MH-Z19C..."); // Takes 3 minutes.
  while(!mhz19c.isReady()) {
    led_off();
    delay(999);
    led_rgb(LOW, LOW, HIGH);
    delay(1);
  }
  Serial.println("MH-Z19C ready");

  // Create the BLE Device
  BLEDevice::init("Nano ESP32");

  // Create the BLE Server
  pServer = BLEDevice::createServer();
  pServer->setCallbacks(new MyServerCallbacks());

  // Create the BLE Service
  BLEService *pService = pServer->createService(SERVICE_UUID);

  // Create a BLE Characteristic
  pTxCharacteristic = pService->createCharacteristic(
										CHARACTERISTIC_UUID_TX,
										BLECharacteristic::PROPERTY_NOTIFY
									);

  pTxCharacteristic->addDescriptor(new BLE2902());

  BLECharacteristic * pRxCharacteristic = pService->createCharacteristic(
											 CHARACTERISTIC_UUID_RX,
											BLECharacteristic::PROPERTY_WRITE
										);

  pRxCharacteristic->setCallbacks(new MyCallbacks());

  // Start the service
  pService->start();

  // Start advertising
  pServer->getAdvertising()->start();

  // LED indicate advertizing
  led_rgb(LOW, LOW, LOW);
}

void loop() {

  if (deviceConnected) {
    mhz19c.getCarbonDioxide(&ppm[0]);
    pTxCharacteristic->setValue(ppm, 2);
    pTxCharacteristic->notify();
    char s[100];
    sprintf(s, "mhz19c: %d ppm [%x %x]", ppm[0]*256+ppm[1], ppm[0], ppm[1]);
    // sprintf(s, "%x %x %x %x %x %x %x %x %x", ppm[0], ppm[1], ppm[2], ppm[3], ppm[4], ppm[5], ppm[6], ppm[7], ppm[8]);
    Serial.println(s);

	delay(500); // bluetooth stack will go into congestion, if too many packets are sent
  }

  // disconnecting
  if (!deviceConnected && oldDeviceConnected) {
    delay(500); // give the bluetooth stack the chance to get things ready
    pServer->startAdvertising(); // restart advertising
    Serial.println("start advertising");
    led_rgb(LOW, LOW, LOW);
    oldDeviceConnected = deviceConnected;
  }

  // connecting
  if (deviceConnected && !oldDeviceConnected) {
    // do stuff here on connecting
    led_rgb(HIGH, HIGH, LOW);
    oldDeviceConnected = deviceConnected;
  }
}
