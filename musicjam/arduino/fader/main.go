package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

var (
	led    = machine.LED
	button = machine.D12
	fader  = machine.ADC{machine.A1}

	note = midi.C3

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	pressed = false
)

func main() {
	machine.InitADC()

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	fader.Configure(machine.ADCConfig{})

	for {
		switch {
		case press():
			led.High()
			midi.Port().NoteOn(midicable, midichannel, note, velocity)

		case release():
			led.Low()
			midi.Port().NoteOff(midicable, midichannel, note, velocity)
		}

		// fader controls volume
		vol := fader.Get()

		if pressed {
			// scale to range 0x0 thru 0x7f
			svol := 0x7F * uint32(vol) / 0xFFFF
			midi.Port().ControlChange(midicable, midichannel, midi.CCVolume, uint8(svol))
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func press() bool {
	if !button.Get() && !pressed {
		pressed = true
		return true
	}
	return false
}

func release() bool {
	if button.Get() && pressed {
		pressed = false
		return true
	}
	return false
}
