package main

import (
	"machine"
	"machine/usb/adc/midi"

	"time"
)

type chord struct {
	name  string
	notes []midi.Note
}

var (
	led                 = machine.LED
	buttonC machine.Pin = machine.D12
	buttonE machine.Pin = machine.D11
	buttonG machine.Pin = machine.D10
	buttonB machine.Pin = machine.D9

	chords = []chord{
		{name: "C ", notes: []midi.Note{midi.C3, midi.E3, midi.G3}},
		{name: "G ", notes: []midi.Note{midi.G3, midi.B3, midi.D4}},
		{name: "Am", notes: []midi.Note{midi.A3, midi.C4, midi.E4}},
		{name: "F ", notes: []midi.Note{midi.F3, midi.A3, midi.C4}},
	}

	keys = []key{
		{name: "C", pin: buttonC, chord: chords[0]},
		{name: "G", pin: buttonE, chord: chords[1]},
		{name: "A", pin: buttonG, chord: chords[2]},
		{name: "F", pin: buttonB, chord: chords[3]},
	}

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	// change this to speed up or slow down the music
	bpm int = 80

	currentNote  int = 0
	currentNotes []midi.Note
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttonC.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonE.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonG.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	go handleKeys()

	for {
		playArpeggio()
		time.Sleep(time.Duration(60000/bpm/4) * time.Millisecond)
	}
}

func handleKeys() {
	for {
		for i := range keys {
			switch {
			case keys[i].press():
				led.High()

				startArpeggio(keys[i].chord.notes)

				keys[i].pressed = true

			case keys[i].release():
				led.Low()

				stopArpeggio()

				keys[i].pressed = false
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

type key struct {
	name    string
	pin     machine.Pin
	chord   chord
	pressed bool
}

func (k key) press() bool {
	if !k.pin.Get() && !k.pressed {
		return true
	}
	return false
}

func (k key) release() bool {
	if k.pin.Get() && k.pressed {
		return true
	}
	return false
}

func startArpeggio(notes []midi.Note) {
	if currentNotes != nil {
		stopArpeggio()
	}

	currentNotes = notes

	midi.Port().NoteOn(midicable, midichannel, currentNotes[currentNote], velocity)
}

func stopArpeggio() {
	if currentNotes == nil {
		return
	}

	midi.Port().NoteOff(midicable, midichannel, currentNotes[currentNote], velocity)

	currentNotes = nil
}

func playArpeggio() {
	if currentNotes == nil {
		return
	}

	midi.Port().NoteOff(midicable, midichannel, currentNotes[currentNote], velocity)

	currentNote++
	if currentNote >= len(currentNotes) {
		currentNote = 0
	}

	midi.Port().NoteOn(midicable, midichannel, currentNotes[currentNote], velocity)
}
