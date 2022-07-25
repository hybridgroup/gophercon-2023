package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

var (
	blue = machine.D12
	green = machine.D10
	button = machine.D11
	touch = machine.D9
	bzrPin = machine.D8
)

func main() {
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	touch.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	for {
		if button.Get() {
			blue.High()
			green.Low()
		} else {
			blue.Low()
			green.High()
		}

		if touch.Get() {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 100)
	}
}
