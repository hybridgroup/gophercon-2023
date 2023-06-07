package main

import (
	"fmt"
	"os"
	"time"

	gobot "gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/sphero/ollie"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := ollie.NewDriver(bleAdaptor)

	rover.On("collision", func(data interface{}) {
		fmt.Printf("collision detected = %+v \n", data)
		rover.SetRGB(255, 0, 0)
	})

	work := func() {
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

			time.Sleep(100 * time.Millisecond)
		}
	}

	robot := gobot.NewRobot("rover",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{rover},
		work,
	)

	robot.Start()
}
