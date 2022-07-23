package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.D12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if button.Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
