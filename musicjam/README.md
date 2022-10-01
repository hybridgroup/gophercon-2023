# TinyGo Music Jam

Make music using your own Arduino-based customized MIDI controller using audio software running on your notebook computer.

```
┌────────────────────────────┐      ┌────────────────────────────────────────────────┐
│                            │      │                                                │
│ ┌────────────────────────┐ │      │ ┌──────────────────────┐                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ │     MIDI Controller    │ │      │ │       USB-MIDI       │                       │
│ │                        ├─┼──────┼─►                      │                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ └────────────────────────┘ │      │ └──────────┬───────────┘                       │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │ ┌──────────▼───────────┐                       │
│                            │      │ │                      ├─────────────────────┐ │
│                            │      │ │                      │                     │ │
│                            │      │ │     Web MIDI API     │ Web Software Synth  │ │
│                            │      │ │                      │                     │ │
│                            │      │ │                      ├─────────────────────┘ │
│                            │      │ │                      │                       │
│                            │      │ │                      │                       │
│                            │      │ │                      │                       │
│                            │      │ └──────────────────────┘                       │
│                            │      │                                                │
└────────────────────────────┘      └────────────────────────────────────────────────┘

  Arduino                             Computer

```

Thanks to USB-MIDI standard, the Arduino will appear as a standard MIDI controller. You can use it to connect to online instruments that use the Web MIDI API.


## Online Synths and Instruments

This is just a list of a few of the available online synths and other virtual instruments.

https://midi.city/

https://www.websynths.com/microtonal/

https://www.gsn-lib.org/apps/cardboardsynth/index.html

https://www.webaudiomodules.org/wamsynths/webcz101/

https://juno-106.js.org/

https://virtualpiano.eu/

https://experiments.withgoogle.com/ai/sound-maker/view/

## Controller

The MIDI Controller is intended to run directly on Arduino to send MIDI commands via the USB interface.

There are several different kinds of controllers in this folder.

We also have some "MakeyButton" boards, which make it possible to become part of the circuit in a very literal sense by using your own body and conductive items to control the music.

### onenote

This introductory MIDI controller sends only a single note. It is designed to connect to some conductive items such as a banana or other piece of fruit.

- Connect one of the "Ground" pins on the Arduino to the breadboard's ground rail (-) using a black or brown jumper cable.

- Connect the "3.3V" pin on the Arduino to the breadboard's power rail (+) using a red jumper cable.

- Connect one of the 10M Ohm resistors from one of the available pins in row 9 on the breadboard, to the breadboard's power rail (+).

- Connect a black cable to the breadboard's ground rail (-). You will hold this cable to serve as the ground.

- Connect a red cable to one of the other available pins in row 9 on the breadboard, and then plug in into your piece of fruit.

- Connect a yellow cable from one of the other available pins in row 9 on the breadboard to pin D12 on the Arduino.

To build/flash the `onenote` example on Arduino Nano RP2040:

        tinygo flash -target nano-rp2040 ./onenote/

Touch the black ground cable, and then also touch the piece of fruit.

This should send MIDI messages tht can trigger sounds on your computer by using your Arduino MIDI controller.

Launch one of the online synths. You should be able to use your new custom MIDI controller to make music.

Have fun!
