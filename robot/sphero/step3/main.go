package main

import (
	"fmt"
	"os"
	"time"

	gobot "gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/sphero/ollie"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := ollie.NewDriver(bleAdaptor)

	work := func() {
		rover.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			rover.SetRGB(255, 0, 0)
		})

		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			rover.SetRGB(r, g, b)
		})

		gobot.Every(3*time.Second, func() {
			rover.Roll(40, uint16(gobot.Rand(360)))
		})
	}

	robot := gobot.NewRobot("rover",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{rover},
		work,
	)

	robot.Start()
}
