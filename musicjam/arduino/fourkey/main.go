package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

var (
	led                 = machine.LED
	buttonC machine.Pin = machine.D12
	buttonE machine.Pin = machine.D11
	buttonG machine.Pin = machine.D10
	buttonB machine.Pin = machine.D9

	keys = []key{
		{pin: buttonC, note: midi.C3},
		{pin: buttonE, note: midi.E3},
		{pin: buttonG, note: midi.G3},
		{pin: buttonB, note: midi.B3},
	}

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttonC.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonE.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonG.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		handleKeys()
		time.Sleep(100 * time.Millisecond)
	}
}

func handleKeys() {
	for i := range keys {
		switch {
		case keys[i].press():
			led.High()
			midi.Port().NoteOn(midicable, midichannel, keys[i].note, velocity)
			keys[i].pressed = true

		case keys[i].release():
			led.Low()
			midi.Port().NoteOff(midicable, midichannel, keys[i].note, velocity)
			keys[i].pressed = false
		}
	}
}

type key struct {
	pin     machine.Pin
	note    midi.Note
	pressed bool
}

func (k key) press() bool {
	if !k.pin.Get() && !k.pressed {
		return true
	}
	return false
}

func (k key) release() bool {
	if k.pin.Get() && k.pressed {
		return true
	}
	return false
}
