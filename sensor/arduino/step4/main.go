package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

var (
	green = machine.D12
	red = machine.D10
	button = machine.D11
	touch = machine.D9
	bzrPin = machine.D8
)

func main() {
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	touch.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	for {
		if button.Get() {
			green.High()
			red.Low()
		} else {
			green.Low()
			red.High()
		}

		if touch.Get() {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 100)
	}
}
