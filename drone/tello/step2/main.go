package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2/platforms/dji/tello"
)

var currentFlightData *tello.FlightData

func main() {
	drone := tello.NewDriver("8888")
	drone.On(tello.FlightDataEvent, func(data interface{}) {
		fd := data.(*tello.FlightData)
		currentFlightData = fd
	})

	drone.Start()
	time.Sleep(2 * time.Second)

	fmt.Println("takeoff...")
	drone.TakeOff()

	start := time.Now()
	for {
		printFlightData(currentFlightData)

		if time.Since(start) > 10*time.Second {
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
