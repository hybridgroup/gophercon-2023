package main

import (
	"machine"
	"time"
)

var (
	blue = machine.D12
	green = machine.D10
	button = machine.D11
)

func main() {
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if button.Get() {
			blue.High()
			green.Low()
		} else {
			blue.Low()
			green.High()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
