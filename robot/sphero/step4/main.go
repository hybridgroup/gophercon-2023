package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/keyboard"
	"gobot.io/x/gobot/platforms/sphero/ollie"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := ollie.NewDriver(bleAdaptor)

	keys := keyboard.NewDriver()

	work := func() {
		rover.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			rover.SetRGB(255, 0, 0)
		})

		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			switch key.Key {
			case keyboard.ArrowUp:
				rover.Roll(40, 0)
			case keyboard.ArrowRight:
				rover.Roll(40, 90)
			case keyboard.ArrowDown:
				rover.Roll(40, 180)
			case keyboard.ArrowLeft:
				rover.Roll(40, 270)
			case keyboard.Spacebar:
				rover.Stop()
			}
		})
	}

	robot := gobot.NewRobot("rover",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{rover, keys},
		work,
	)

	robot.Start()
}
