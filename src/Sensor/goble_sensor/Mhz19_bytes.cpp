#include "Mhz19_bytes.h"
#include <assert.h>
#include <cstdint>

const uint8_t Mhz19c::CommandRead[Mhz19c::PacketLength] = {
    0xFF, 0x01, 0x86, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79};
const uint8_t Mhz19c::CommandEnableAutoBaseCalibration[Mhz19c::PacketLength] = {
    0xFF, 0x01, 0x79, 0xA0, 0x00, 0x00, 0x00, 0x00, 0xE6};
const uint8_t Mhz19c::CommandDisableAutoBaseCalibration[Mhz19c::PacketLength] = {
    0xFF, 0x01, 0x79, 0x00, 0x00, 0x00, 0x00, 0x00, 0x86};
const uint8_t Mhz19c::CommandCalibrateToZeroPoint[Mhz19c::PacketLength] = {
    0xFF, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x78};

Mhz19c::Mhz19c() : serial_(nullptr), isPreheatingDone_(false) {}

#ifdef ARDUINO_MHZ19_UNIT_TEST

Mhz19c::~Mhz19c() {}

#endif

void Mhz19c::begin(Stream* serial) { serial_ = serial; }

bool Mhz19c::isReady() const {
  if (isPreheatingDone_) {
    return true;
  }

  if (millis() > MHZ19_PREHEATING_DURATION) {
    isPreheatingDone_ = true;
    return true;
  }

  return false;
}

void Mhz19c::getCarbonDioxide(uint8_t *result) const {
  assert(serial_ != nullptr);

  if (!isReady()) {
    return;
  }

  uint8_t response[PacketLength];
  serial_->write(CommandRead, PacketLength);

  if (serial_->available()) {
    serial_->readBytes(response, PacketLength);

    // result[0] = response[2];
    // result[1] = response[3];
    // result[2] = response[8];

    auto checkSum = calculatePacketCheckSum(response);
    if (response[0] == 0xFF && response[1] == CommandRead[2] &&
        response[8] == checkSum) {
      result[0] = response[2];
      result[1] = response[3];
    } else {
      result[0] = 0x01;
      result[1] = 0x01;
    }
  } else {
      result[0] = 0x00;
      result[1] = 0x00;
  }

  return;
}

bool Mhz19c::setMeasuringRange(const Mhz19MeasuringRange measuringRange) {
  auto low = static_cast<uint8_t>(static_cast<uint16_t>(measuringRange) % 256);
  auto high = static_cast<uint8_t>(static_cast<uint16_t>(measuringRange) / 256);

  uint8_t command[PacketLength] = {0xFF, 0x01, 0x99, 0x00, 0x00,
                                   0x00, high, low,  0x00};
  command[8] = calculatePacketCheckSum(command);

  return sendCommand(command);
}

bool Mhz19c::enableAutoBaseCalibration() {
  return sendCommand(CommandEnableAutoBaseCalibration);
}

bool Mhz19c::disableAutoBaseCalibration() {
  return sendCommand(CommandDisableAutoBaseCalibration);
}

bool Mhz19c::calibrateToZeroPoint() {
  return sendCommand(CommandCalibrateToZeroPoint);
}

bool Mhz19c::calibrateToSpanPoint(const uint16_t spanPoint) {
  if ((spanPoint < static_cast<uint16_t>(Mhz19MeasuringRange::Ppm_1000)) ||
      (spanPoint > static_cast<uint16_t>(Mhz19MeasuringRange::Ppm_5000))) {
    return false;
  }

  auto low = static_cast<uint8_t>(spanPoint % 256);
  auto high = static_cast<uint8_t>(spanPoint / 256);

  uint8_t command[PacketLength] = {0xFF, 0x01, 0x88, high, low,
                                   0x00, 0x00, 0x00, 0x00};
  command[8] = calculatePacketCheckSum(command);

  return sendCommand(command);
}

uint8_t Mhz19c::calculatePacketCheckSum(const uint8_t* packet) {
  uint8_t checkSum = 0;

  for (size_t i = 1; i < PacketLength - 1; i++) {
    checkSum += packet[i];
  }

  checkSum = 255 - checkSum;
  checkSum++;
  return checkSum;
}

bool Mhz19c::sendCommand(const uint8_t* command) const {
  assert(serial_ != nullptr);

  uint8_t response[PacketLength];

  serial_->write(command, PacketLength);
  serial_->readBytes(response, PacketLength);

  auto checkSum = calculatePacketCheckSum(response);
  if (response[0] != 0xFF || response[1] != command[2] ||
      response[8] != checkSum) {
    return false;
  }

  return true;
}
