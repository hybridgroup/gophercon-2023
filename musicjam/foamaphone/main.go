package main

import (
	"machine"
	"machine/usb/adc/midi"
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
	buttonG machine.Pin = machine.D8

	keys [4]*makeybutton.Button

	notes             = []midi.Note{midi.C4, midi.E4, midi.G4, midi.B4}
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	initKeys()

	for {
		handleKeys()
		time.Sleep(100 * time.Millisecond)
	}
}

func initKeys() {
	keys[0] = makeybutton.NewButton(buttonC)
	keys[1] = makeybutton.NewButton(buttonD)
	keys[2] = makeybutton.NewButton(buttonE)
	keys[3] = makeybutton.NewButton(buttonG)

	for key := 0; key < numberOfKeys; key++ {
		keys[key].Configure()
	}
}

func handleKeys() {
	for key := 0; key < numberOfKeys; key++ {
		switch keys[key].Get() {
		case makeybutton.Pressed:
			midi.Port().NoteOn(0, midichannel, notes[key], 50)
		case makeybutton.Released:
			midi.Port().NoteOff(0, midichannel, notes[key], 50)
		}
	}
}
