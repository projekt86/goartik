package goartik

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	// IN : pin in mode
	IN = PinMode("in")
	// OUT : pin out mode
	OUT = PinMode("out")
	// LOW : value 0
	LOW = PinValue(0)
	// HIGH : value 1
	HIGH = PinValue(1)
)

const (
	exportTestPath    = "/test/device/sys/class/gpio/export"
	pinTestPath       = "/test/device/sys/class/gpio/gpio%d"
	directionTestPath = "/test/device/sys/class/gpio/gpio%d/direction"
	valueTestPath     = "/test/device/sys/class/gpio/gpio%d/value"

	exportPath    = "/sys/class/gpio/export"
	pinPath       = "/sys/class/gpio/gpio%d"
	directionPath = "/sys/class/gpio/gpio%d/direction"
	valuePath     = "/sys/class/gpio/gpio%d/value"

	isVirtualDevice = false
)

// PinMode : pin mode
type PinMode string

// PinValue : pin value
type PinValue uint8

func (v PinValue) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

// DigitalPinMode : set pin mode
func (m *ArtikModule) DigitalPinMode(pin uint, mode PinMode) error {
	pin = m.mapPin(pin)

	err := m.exportPin(pin)
	if err != nil {
		fmt.Println("Can't export pin")
		return err
	}

	var path string
	if isVirtualDevice {
		path = getRelativePath(fmt.Sprintf(directionTestPath, pin))
	} else {
		path = fmt.Sprintf(directionPath, pin)
	}

	err = ioutil.WriteFile(path, []byte(mode), 0644)
	if err != nil {
		fmt.Println("Can't set pin mode")
		return err
	}

	return nil
}

// DigitalRead : read value of pin
func (m *ArtikModule) DigitalRead(pin uint) (PinValue, error) {
	pin = m.mapPin(pin)

	var path string
	if isVirtualDevice {
		path = getRelativePath(fmt.Sprintf(valueTestPath, pin))
	} else {
		path = fmt.Sprintf(valuePath, pin)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return convPinValue(b), nil
}

// DigitalWrite : write value to pin
func (m *ArtikModule) DigitalWrite(pin uint, value PinValue) error {
	pin = m.mapPin(pin)

	var path string
	if isVirtualDevice {
		path = getRelativePath(fmt.Sprintf(valueTestPath, pin))
	} else {
		path = fmt.Sprintf(valuePath, pin)
	}

	err := ioutil.WriteFile(path, []byte(value.String()), 0664)
	if err != nil {
		return err
	}

	return nil
}

func (m *ArtikModule) mapPin(pin uint) uint {
	p := strconv.FormatUint(uint64(pin), 10)
	mappedPin, ok := m.config.Pins[p]
	if ok {
		return mappedPin
	}
	return pin
}

func (m *ArtikModule) exportPin(pin uint) error {
	pin = m.mapPin(pin)

	if isVirtualDevice {
		path := getRelativePath(fmt.Sprintf(pinTestPath, pin))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, 0777); err == nil {
				ioutil.WriteFile(fmt.Sprintf("%s/direction", path), []byte{}, 0664)
				ioutil.WriteFile(fmt.Sprintf("%s/value", path), []byte("0"), 0664)
			}
		}
		return nil
	}

	s := strconv.FormatUint(uint64(pin), 10)
	return ioutil.WriteFile(exportPath, []byte(s), 0644)
}

func convPinValue(b []byte) PinValue {
	v, _ := strconv.Atoi(string(b))
	return PinValue(v)
}
