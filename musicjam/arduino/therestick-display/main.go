package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

var (
	led    = machine.LED
	button = machine.D12
	stickX = machine.ADC{machine.A2}
	stickY = machine.ADC{machine.A3}

	note = midi.C3

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	sx, sy  uint32
	pressed = false
)

func main() {
	machine.InitADC()

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	stickX.Configure(machine.ADCConfig{})
	stickY.Configure(machine.ADCConfig{})

	go handleDisplay()

	for {
		switch {
		case press():
			led.High()
			midi.Port().NoteOn(midicable, midichannel, note, velocity)

		case release():
			led.Low()
			midi.Port().NoteOff(midicable, midichannel, note, velocity)

			sx, sy = 0, 0
			midi.Port().ControlChange(midicable, midichannel, midi.CCModulationWheel, 0)
			midi.Port().PitchBend(midicable, midichannel, 0)
		}

		if pressed {
			// x axis for modulation, y axis for pitch bend
			x, y := stickX.Get(), stickY.Get()

			// scale x to range 0x0 thru 0xff
			sx = 0xFF * uint32(x) / 0xFFFF
			midi.Port().ControlChange(midicable, midichannel, midi.CCModulationWheel, uint8(sx))

			// scale y to range 0x0 thru 0x3FFF
			sy = 0x3FFF * uint32(y) / 0xFFFF
			midi.Port().PitchBend(midicable, midichannel, uint16(sy))
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
