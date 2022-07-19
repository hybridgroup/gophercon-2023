package main

import (
	"machine"
	"time"
)

func main() {
	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	green := machine.D10
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !button.Get() {
			blue.Low()
			green.High()
		} else {
			blue.High()
			green.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
