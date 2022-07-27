package main

import (
	"machine"
	"machine/usb/midi"
	"time"

	"tinygo.org/x/drivers/makeybutton"
)

const (
	numberOfKeys = 4
)

var (
	buttonC machine.Pin = machine.D2
	buttonD machine.Pin = machine.D4
	buttonE machine.Pin = machine.D6
	buttonG machine.Pin = machine.D9

	keys        [4]*makeybutton.Button
	notes       = []midi.Note{midi.C4, midi.E4, midi.G4, midi.B4}
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	initKeys()

	for {
		readKeys()
		time.Sleep(100 * time.Millisecond)
	}
}

func readKeys() {
	for key := 0; key < numberOfKeys; key++ {
		switch keys[key].Get() {
		case makeybutton.Pressed:
			midi.Midi.NoteOn(0, midichannel, notes[key], 50)
		case makeybutton.Released:
			midi.Midi.NoteOff(0, midichannel, notes[key], 50)
		}
	}
}

func initKeys() {
	keys[0] = makeybutton.NewButton(buttonC)
	keys[1] = makeybutton.NewButton(buttonD)
	keys[2] = makeybutton.NewButton(buttonE)
	keys[3] = makeybutton.NewButton(buttonG)
}
