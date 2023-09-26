package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"

	"tinygo.org/x/drivers/shifter"
)

var (
	led     = machine.LED
	buttons shifter.Device

	buttonC = shifter.BUTTON_UP
	buttonE = shifter.BUTTON_RIGHT
	buttonG = shifter.BUTTON_DOWN
	buttonB = shifter.BUTTON_LEFT

	keys = []key{
		{name: "C", pin: buttonC, note: midi.C3},
		{name: "E", pin: buttonE, note: midi.E3},
		{name: "G", pin: buttonG, note: midi.G3},
		{name: "B", pin: buttonB, note: midi.B3},
	}

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttons = shifter.NewButtons()
	buttons.Configure()

	go handleDisplay()

	for {
		handleKeys()
		time.Sleep(100 * time.Millisecond)
	}
}

func handleKeys() {
	buttons.ReadInput()

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
	name    string
	pin     int
	note    midi.Note
	pressed bool
}

func (k key) press() bool {
	if buttons.Pins[k.pin].Get() && !k.pressed {
		return true
	}
	return false
}

func (k key) release() bool {
	if !buttons.Pins[k.pin].Get() && k.pressed {
		return true
	}
	return false
}
