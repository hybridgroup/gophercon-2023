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

	stickX = machine.ADC{machine.A8}
	stickY = machine.ADC{machine.A9}

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

	buttons = shifter.NewButtons()
	buttons.Configure()

	stickX.Configure(machine.ADCConfig{})
	stickY.Configure(machine.ADCConfig{})

	go handleDisplay()

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

		// x axis for modulation, y axis for pitch bend
		x, y := stickX.Get(), stickY.Get()

		if pressed {
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
