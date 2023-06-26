package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	gobot "gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/mqtt"
	"gobot.io/x/gobot/v2/platforms/sphero/ollie"

	term "github.com/nsf/termbox-go"
)

var robot *gobot.Robot
var mqttAdaptor *mqtt.Adaptor

const mqtttopic = "tinygo/hacksession/heartbeat"

func ReportCollision(data interface{}) {
	buf := new(bytes.Buffer)
	msg, _ := json.Marshal(data)
	binary.Write(buf, binary.LittleEndian, msg)
	mqttAdaptor.Publish("tinygo/hacksession/collision/"+robot.Name, buf.Bytes())
}

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := ollie.NewDriver(bleAdaptor)

	mqttAdaptor = mqtt.NewAdaptor(os.Args[2], "rover")
	mqttAdaptor.SetAutoReconnect(true)

	heartbeat := mqtt.NewDriver(mqttAdaptor, mqtttopic)

	work := func() {
		rover.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			rover.SetRGB(255, 0, 0)
			ReportCollision(data)
		})

		heartbeat.On(mqtt.Data, func(data interface{}) {
			fmt.Println("heartbeat")
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			rover.SetRGB(r, g, b)
		})

		defer term.Close()

		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyEsc:
					term.Sync()
					fmt.Println("exiting...")
					return
				default:
					term.Sync()
					// WSAD to control forward, backward, left, and right
					switch ev.Ch {
					case 'w':
						fmt.Println("forward...")
						rover.Roll(60, 0)
					case 's':
						fmt.Println("backward...")
						rover.Roll(60, 180)
					case 'a':
						fmt.Println("left...")
						rover.Roll(60, 270)
					case 'd':
						fmt.Println("right...")
						rover.Roll(60, 90)
					default:
						rover.Stop()
					}
				}
			case term.EventError:
				panic(ev.Err)
			}
		}
	}

	robot = gobot.NewRobot([]gobot.Connection{bleAdaptor, mqttAdaptor},
		[]gobot.Device{rover},
		work,
	)

	robot.Start()
}
