package main

import (
	"fmt"
	"os"
	"time"

	gobot "gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/ble"
	"gobot.io/x/gobot/v2/platforms/parrot/minidrone"

	term "github.com/nsf/termbox-go"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	bleAdaptor.Connect()

	drone := minidrone.NewDriver(bleAdaptor)
	drone.Start()

	err := term.Init()
	if err != nil {
		panic(err)
	}

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
					case '[':
						fmt.Println("takeoff...")
						drone.TakeOff()
					case ']':
						fmt.Println("land...")
						drone.Land()
					case 'w':
						fmt.Println("forward...")
						drone.Forward(20)
					case 's':
						fmt.Println("backward...")
						drone.Backward(20)
					case 'a':
						fmt.Println("left...")
						drone.Left(20)
					case 'd':
						fmt.Println("right...")
						drone.Right(20)
					// IKJL to control up, down, spin counter clockwise, spin clockwise
					case 'i':
						fmt.Println("up...")
						drone.Up(20)
					case 'k':
						fmt.Println("down...")
						drone.Down(20)
					case 'j':
						fmt.Println("spin counter clockwise...")
						drone.CounterClockwise(minidrone.ValidatePitch(20, 10))
					case 'l':
						fmt.Println("spin clockwise...")
						drone.Clockwise(minidrone.ValidatePitch(20, 10))
					// TGFH to flip front, flip back, flip left, flip right
					case 't':
						fmt.Println("front flip...")
						drone.FrontFlip()
					case 'g':
						fmt.Println("back flip...")
						drone.BackFlip()
					case 'f':
						fmt.Println("left flip...")
						drone.LeftFlip()
					case 'h':
						fmt.Println("right flip...")
						drone.RightFlip()
					default:
						drone.Stop()
					}
				}
			case term.EventError:
				panic(ev.Err)
			}

			time.Sleep(100 * time.Millisecond)
		}
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
