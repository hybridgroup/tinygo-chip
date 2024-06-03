package main

import (
	"time"

	chip "github.com/hybridgroup/tinygo-chip"
	"tinygo.org/x/bluetooth"
)

// replace this with the MAC address of the Bluetooth peripheral you want to connect to.
const deviceAddress = "1C:1D:53:27:15:2E"

var (
	adapter = bluetooth.DefaultAdapter
	device  bluetooth.Device
	ch      = make(chan bluetooth.ScanResult, 1)

	robot *chip.Robot
)

func main() {
	time.Sleep(5 * time.Second)
	println("enabling...")

	must("enable BLE interface", adapter.Enable())

	println("start scan...")

	must("start scan", adapter.Scan(scanHandler))

	var err error
	select {
	case result := <-ch:
		device, err = adapter.Connect(result.Address, bluetooth.ConnectionParams{})
		must("connect to peripheral device", err)

		println("connected to ", result.Address.String())
	}

	defer device.Disconnect()

	robot = chip.NewRobot(&device)
	err = robot.Start()
	if err != nil {
		println(err)
	}

	println("lie down...")
	err = robot.Action(chip.ActionLieDown)
	if err != nil {
		println(err)
	}

	time.Sleep(3 * time.Second)

	println("dance...")
	err = robot.Action(chip.ActionDance)
	if err != nil {
		println(err)
	}

	time.Sleep(3 * time.Second)

	robot.Action(chip.ActionReset)
}

func scanHandler(a *bluetooth.Adapter, d bluetooth.ScanResult) {
	println("device:", d.Address.String(), d.RSSI, d.LocalName())
	if d.Address.String() == deviceAddress {
		a.StopScan()
		ch <- d
	}
}

func must(action string, err error) {
	if err != nil {
		for {
			println("failed to " + action + ": " + err.Error())
			time.Sleep(time.Second)
		}
	}
}
