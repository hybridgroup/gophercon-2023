package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

var (
	led    = machine.LED
	button = machine.D12

	note = midi.C3

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if button.Get() {
			led.Low()
			midi.Port().NoteOff(midicable, midichannel, note, velocity)
		} else {
			led.High()
			midi.Port().NoteOn(midicable, midichannel, note, velocity)
		}

		time.Sleep(time.Millisecond * 100)
	}
}
