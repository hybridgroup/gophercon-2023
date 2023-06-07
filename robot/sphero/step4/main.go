// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"gobot.io/x/gobot"
// 	"gobot.io/x/gobot/platforms/ble"
// 	"gobot.io/x/gobot/platforms/keyboard"
// 	"gobot.io/x/gobot/platforms/sphero/ollie"
// )

// func main() {
// 	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
// 	rover := ollie.NewDriver(bleAdaptor)

// 	keys := keyboard.NewDriver()

// 	work := func() {
// 		rover.On("collision", func(data interface{}) {
// 			fmt.Printf("collision detected = %+v \n", data)
// 			rover.SetRGB(255, 0, 0)
// 		})

// 		keys.On(keyboard.Key, func(data interface{}) {
// 			key := data.(keyboard.KeyEvent)

// 			switch key.Key {
// 			case keyboard.W:
// 				rover.Roll(60, 0)
// 			case keyboard.D:
// 				rover.Roll(60, 90)
// 			case keyboard.S:
// 				rover.Roll(60, 180)
// 			case keyboard.A:
// 				rover.Roll(60, 270)
// 			case keyboard.Spacebar:
// 				rover.Stop()
// 			}

// 			gobot.After(100 * time.Millisecond, func() {
// 				rover.Stop()
// 			})
// 		})
// 	}

// 	robot := gobot.NewRobot("rover",
// 		[]gobot.Connection{bleAdaptor},
// 		[]gobot.Device{rover, keys},
// 		work,
// 	)

//		robot.Start()
//	}
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
