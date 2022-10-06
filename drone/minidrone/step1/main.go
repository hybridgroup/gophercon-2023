package main

import (
	"fmt"
	"time"
	"os"

	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone := minidrone.NewDriver(bleAdaptor)

	bleAdaptor.Connect()
	drone.Start()

	time.Sleep(2 * time.Second)
	fmt.Println("takeoff...")
	drone.TakeOff()

	start := time.Now()
	for {
		if time.Since(start) > 10*time.Second {
			fmt.Println("landing...")
			drone.Land()
			return
		}

		time.Sleep(time.Second)
	}
}

// package main

// import (
// 	"os"
// 	"time"

// 	minidrone "github.com/hybridgroup/tinygo-minidrone"
// 	"tinygo.org/x/bluetooth"
// )

// // replace this with the MAC address of the Bluetooth peripheral you want to connect to.
// // const deviceAddress = "E0:14:DC:85:3D:D1"

// var address string

// var (
// 	adapter = bluetooth.DefaultAdapter
// 	device  *bluetooth.Device
// 	ch      = make(chan bluetooth.ScanResult, 1)
// 	buf     = make([]byte, 255)

// 	drone *minidrone.Minidrone
// )

// func main() {
// 	address = os.Args[1]

// 	time.Sleep(5 * time.Second)
// 	println("enabling...")

// 	must("enable BLE interface", adapter.Enable())

// 	println("start scan...")

// 	must("start scan", adapter.Scan(scanHandler))

// 	var err error
// 	select {
// 	case result := <-ch:
// 		device, err = adapter.Connect(result.Address, bluetooth.ConnectionParams{})
// 		must("connect to peripheral device", err)

// 		println("connected to ", result.Address.String())
// 	}

// 	defer device.Disconnect()

// 	drone = minidrone.NewMinidrone(device)
// 	err = drone.Start()
// 	if err != nil {
// 		println(err)
// 	}

// 	time.Sleep(3 * time.Second)

// 	println("takeoff")
// 	err = drone.TakeOff()
// 	if err != nil {
// 		println(err)
// 	}
// 	time.Sleep(3 * time.Second)

// 	println("land")
// 	err = drone.Land()
// 	if err != nil {
// 		println(err)
// 	}

// 	drone.Halt()
// }

// func scanHandler(a *bluetooth.Adapter, d bluetooth.ScanResult) {
// 	println("device:", d.Address.String(), d.RSSI, d.LocalName())
// 	if d.Address.String() == address {
// 		a.StopScan()
// 		ch <- d
// 	}
// }

// func must(action string, err error) {
// 	if err != nil {
// 		for {
// 			println("failed to " + action + ": " + err.Error())
// 			time.Sleep(time.Second)
// 		}
// 	}
// }
