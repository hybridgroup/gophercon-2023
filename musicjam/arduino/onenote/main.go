package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

var (
	led    = machine.LED
	button = machine.D12
	
	keyOfMusic = midi.C4
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if button.Get() {
			led.Low()
			midi.Port().NoteOff(0, midichannel, keyOfMusic, 0x40)
		} else {
			led.High()
			midi.Port().NoteOn(0, midichannel, keyOfMusic, 0x40)
		}

		time.Sleep(time.Millisecond * 100)
	}
}
