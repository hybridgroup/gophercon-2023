package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

func main() {
	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	green := machine.D10
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	touch := machine.D9
	touch.Configure(machine.PinConfig{Mode: machine.PinInput})

	bzrPin := machine.D8
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	for {
		if !button.Get() {
			blue.Low()
			green.High()
		} else {
			blue.High()
			green.Low()
		}

		if touch.Get() {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
