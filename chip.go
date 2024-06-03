package chip

import (
	"errors"

	"tinygo.org/x/bluetooth"
)

type Robot struct {
	device        *bluetooth.Device
	receive       *bluetooth.DeviceService
	receiveNotify *bluetooth.DeviceCharacteristic
	send          *bluetooth.DeviceService
	sendData      *bluetooth.DeviceCharacteristic

	buf []byte
}

var (
	// BLE services
	chipReceiveDataService              = bluetooth.New16BitUUID(0xffe0)
	chipReceiveDataNotifyCharacteristic = bluetooth.New16BitUUID(0xffe4)

	chipSendDataService             = bluetooth.New16BitUUID(0xffe5)
	chipSendDataWriteCharacteristic = bluetooth.New16BitUUID(0xffe9)
)

// NewRobot creates a new CHiP robot.
func NewRobot(dev *bluetooth.Device) *Robot {
	r := &Robot{
		device: dev,
		buf:    make([]byte, 255),
	}

	return r
}

func (r *Robot) Start() (err error) {
	srvcs, err := r.device.DiscoverServices([]bluetooth.UUID{
		chipReceiveDataService,
		chipSendDataService,
	})
	if err != nil || len(srvcs) == 0 {
		return errors.New("could not find services")
	}

	r.receive = &srvcs[0]
	r.send = &srvcs[1]
	println("found chip receive service", r.receive.UUID().String())
	println("found chip send service", r.send.UUID().String())

	chars, err := r.receive.DiscoverCharacteristics([]bluetooth.UUID{
		chipReceiveDataNotifyCharacteristic,
	})
	if err != nil || len(chars) == 0 {
		return errors.New("could not find chip receive characteristic")
	}

	r.receiveNotify = &chars[0]

	chars, err = r.send.DiscoverCharacteristics([]bluetooth.UUID{
		chipSendDataWriteCharacteristic,
	})
	if err != nil || len(chars) == 0 {
		return errors.New("could not find chip write characteristic")
	}

	r.sendData = &chars[0]

	return
}

// Action performs one of the actions of the Chip robot.
func (r *Robot) Action(action ActionType) (err error) {
	buf := []byte{Action, uint8(action)}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return
}
