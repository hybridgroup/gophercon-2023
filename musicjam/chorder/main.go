package main

import (
	"machine"
	"machine/usb/adc/midi"
	"time"

	"tinygo.org/x/drivers/makeybutton"
)

const (
	keyOfMusic = midi.C4
)

var (
	buttonC     machine.Pin = machine.D12
	midichannel uint8       = 0 // MIDI channels are 0-15 e.g. 1-16

	chords = []struct {
		name string
		keys []midi.Note
	}{
		{name: "C ", keys: []midi.Note{midi.C4, midi.E4, midi.G4}},
		{name: "G ", keys: []midi.Note{midi.G3, midi.B3, midi.D4}},
		{name: "Am", keys: []midi.Note{midi.A3, midi.C4, midi.E4}},
		{name: "F ", keys: []midi.Note{midi.F3, midi.A3, midi.C4}},
	}
)

func main() {
	key := makeybutton.NewButton(buttonC)
	key.Configure()

	index := 0

	for {
		switch key.Get() {
		case makeybutton.Pressed:
			for _, c := range chords[index].keys {
				midi.Port().NoteOn(0, 0, c, 0x40)
			}
		case makeybutton.Released:
			for _, c := range chords[index].keys {
				midi.Port().NoteOff(0, 0, c, 0x40)
			}
			index = (index + 1) % len(chords)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
