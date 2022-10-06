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
	bleAdaptor.Connect()

	drone := minidrone.NewDriver(bleAdaptor)
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
