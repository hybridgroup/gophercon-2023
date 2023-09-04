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

	note = midi.C3

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	pressed = false
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttons = shifter.NewButtons()
	buttons.Configure()

	for {
		buttons.ReadInput()

		switch {
		case press():
			led.High()
			midi.Port().NoteOn(midicable, midichannel, note, velocity)

		case release():
			led.Low()
			midi.Port().NoteOff(midicable, midichannel, note, velocity)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func press() bool {
	if buttons.Pins[shifter.BUTTON_A].Get() && !pressed {
		pressed = true
		return true
	}
	return false
}

func release() bool {
	if !buttons.Pins[shifter.BUTTON_A].Get() && pressed {
		pressed = false
		return true
	}
	return false
}
