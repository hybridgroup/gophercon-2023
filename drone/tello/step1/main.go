package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	drone.Start()

	time.Sleep(2 * time.Second)
	fmt.Println("takeoff...")
	drone.TakeOff()

	start := time.Now()
	for {
		if time.Since(start) > 10*time.Second {
			fmt.Println("landing...")
			drone.Land()
			return
		}

		time.Sleep(time.Second)
	}
}
