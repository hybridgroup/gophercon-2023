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

	chords = []struct {
		name string
		notes []midi.Note
	}{
		{name: "C ", notes: []midi.Note{midi.C4, midi.E4, midi.G4}},
		{name: "G ", notes: []midi.Note{midi.G3, midi.B3, midi.D4}},
		{name: "Am", notes: []midi.Note{midi.A3, midi.C4, midi.E4}},
		{name: "F ", notes: []midi.Note{midi.F3, midi.A3, midi.C4}},
	}

	pressed bool
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	index := 0

	for {
		switch {
		case press():
			led.High()
			for _, c := range chords[index].notes {
				midi.Port().NoteOn(0, 0, c, 0x40)
			}
		case release():
			led.Low()
			for _, c := range chords[index].notes {
				midi.Port().NoteOff(0, 0, c, 0x40)
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
