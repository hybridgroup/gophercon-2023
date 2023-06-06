package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/parrot/minidrone"
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
