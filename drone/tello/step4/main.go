package main

import (
	"fmt"
	"sync"
	"time"

	"gobot.io/x/gobot/v2/platforms/dji/tello"
)

var (
	drone             = tello.NewDriver("8888")
	currentFlightData *tello.FlightData
	flightPattern     sync.Once
)

func main() {
	drone.On(tello.FlightDataEvent, func(data interface{}) {
		fd := data.(*tello.FlightData)
		currentFlightData = fd
	})
	drone.On(tello.TakeoffEvent, func(data interface{}) {
		fmt.Println("takeoff...")
	})
	drone.On(tello.FlipEvent, func(data interface{}) {
		fmt.Println("Flip!")
	})

	drone.Start()
	time.Sleep(2 * time.Second)

	fmt.Println("Prepare for takeoff...")
	drone.TakeOff()

	start := time.Now()
	for {
		printFlightData(currentFlightData)

		if time.Since(start) > 2*time.Second {
			flightPattern.Do(func() {
				go flyWithFlip()
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

func printFlightData(d *tello.FlightData) {
	if d.BatteryLow {
		fmt.Printf(" -- Battery low: %d%% --\n", d.BatteryPercentage)
	}

	displayData := `
Battery:		%d%%
Height:         %d
Ground Speed:   %d
Light Strength: %d

`
	fmt.Printf(displayData, d.BatteryPercentage, d.Height, d.GroundSpeed, d.LightStrength)
}

func flyWithFlips() {
	drone.Forward(20)
	time.Sleep(time.Second * 3)
	drone.Forward(0)
	drone.Backward(20)
	time.Sleep(time.Second * 3)
	drone.Backward(0)

	performFlips()

	drone.Land()
}

func performFlips() {
	drone.FrontFlip()
	time.Sleep(time.Second * 3)
	drone.BackFlip()
}
