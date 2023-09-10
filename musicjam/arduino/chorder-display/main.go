package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

var (
	led    = machine.LED
	button = machine.D12

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	chords = []struct {
		name  string
		notes []midi.Note
	}{
		{name: "C ", notes: []midi.Note{midi.C3, midi.E3, midi.G3}},
		{name: "G ", notes: []midi.Note{midi.G3, midi.B3, midi.D4}},
		{name: "Am", notes: []midi.Note{midi.A3, midi.C4, midi.E4}},
		{name: "F ", notes: []midi.Note{midi.F3, midi.A3, midi.C4}},
	}

	index int
	pressed bool
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	go handleDisplay()

	for {
		switch {
		case press():
			led.High()
			for _, note := range chords[index].notes {
				midi.Port().NoteOn(midicable, midichannel, note, velocity)
			}
		case release():
			led.Low()
			for _, note := range chords[index].notes {
				midi.Port().NoteOff(midicable, midichannel, note, velocity)
			}
			index = (index + 1) % len(chords)
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

