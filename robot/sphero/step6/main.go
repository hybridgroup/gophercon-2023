package main

import (
	"fmt"
	"os"

	gobot "gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/keyboard"
	"gobot.io/x/gobot/v2/platforms/mqtt"
	"gobot.io/x/gobot/v2/platforms/sphero/ollie"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := ollie.NewDriver(bleAdaptor)

	mqttAdaptor := mqtt.NewAdaptor(os.Args[2], "rover")
	mqttAdaptor.SetAutoReconnect(true)

	heartbeat := mqtt.NewDriver(mqttAdaptor, "basestation/heartbeat")

	keys := keyboard.NewDriver()

	work := func() {
		rover.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			rover.SetRGB(255, 0, 0)
		})

		heartbeat.On(mqtt.Data, func(data interface{}) {
			fmt.Println("heartbeat")
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			rover.SetRGB(r, g, b)
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
		[]gobot.Connection{bleAdaptor, mqttAdaptor},
		[]gobot.Device{rover, keys},
		work,
	)

	robot.Start()
}
