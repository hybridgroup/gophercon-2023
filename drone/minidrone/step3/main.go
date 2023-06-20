package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/parrot/minidrone"
)

var (
	drone         *minidrone.Driver
	flightPattern sync.Once
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	bleAdaptor.Connect()

	drone = minidrone.NewDriver(bleAdaptor)
	drone.Start()

	go func() {
		droneEvents := drone.Subscribe()

		for event := range droneEvents {
			fmt.Println("Event:", event.Name, event.Data)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("takeoff...")
	drone.TakeOff()

	start := time.Now()
	for {
		if time.Since(start) > 5*time.Second {
			flightPattern.Do(func() {
				go flySimpleMovements()
			})
		}

		if time.Since(start) > 20*time.Second {
			fmt.Println("landing...")
			drone.Land()
			return
		}

		time.Sleep(time.Second)
	}
}

func flySimpleMovements() {
	drone.Forward(20)
	time.Sleep(time.Second * 3)
	drone.Forward(0)
	drone.Backward(20)
	time.Sleep(time.Second * 3)
	drone.Backward(0)
	drone.Left(20)
	time.Sleep(time.Second * 3)
	drone.Left(0)
	drone.Right(20)
	time.Sleep(time.Second * 3)
	drone.Right(0)
	drone.Land()
}
