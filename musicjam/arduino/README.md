# MIDI Music Jam using the Arduino RP2040 Nano

## What you need

    - Arduino Nano RP2040 Connect IoT board
    - MIDI Starter Kit parts
    - Personal computer with Go 1.19+ and TinyGo installed, and a USB port.

### TinyGo drivers

All of the code dependencies you will need are already in the Go modules file in this directory, so they will be downloaded and installed automatically. You don't need to do anything, when you follow the subsequent instructions they will be downloaded by TinyGo.

Just for your information, the TinyGo drivers that let you connect to sensors, displays, and other external peripheral devices are located in the separate repository at https://github.com/tinygo-org/drivers


## Connecting the Arduino Nano RP2040 Connect to your computer

<img src="https://docs.arduino.cc/static/8b9e4e17c1e1afa836057c5ba87c27c9/2f891/pinout.png" alt="Arduino Nano RP2040 Connect" width="600"/>

Plug the Arduino Nano RP2040 Connect into your computer using a USB cable. There may be one provided in your starter kit.

## Running the code

The TinyGo programs will run directly on the Arduino microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your Arduino.
- The program executes from the Arduino.

Let's get started!

## Code

### hello - Built-in LED

![Arduino Nano RP2040 Connect](./assets/step0.jpg)

This tests that you can compile and flash your Arduino with TinyGo code, by blinking the built-in LED.

Run the following command to compile your code, and flash it onto the Arduino:

```
tinygo flash -target nano-rp2040 ./hello/
```

Once the Arduino is flashed correctly, the built-in amber LED to the right of the USB jack should start to turn on and off once per second. Now everything is setup correctly and you are ready to continue.

### onenote

This introductory MIDI controller sends only a single note.

- Connect one of the "Ground" pins on the Arduino to the breadboard's ground rail (-).

- Connect a black cable from one pin of the blue button on the breadboard to the breadboard's ground rail (-).

- Connect a colored cable from the other pin of the blue button on the breadboard to pin D12 on the Arduino.

To build/flash the `onenote` example on Arduino Nano RP2040:

        tinygo flash -target nano-rp2040 ./onenote/

Press the button.

This should send MIDI messages that can trigger sounds on your computer by using your Arduino MIDI controller.

Open a web page with one of the online synths listed above. You should be able to use your new custom MIDI controller to make music.

Have fun!

### chorder

This MIDI controller sends entire chords with a single touch. It uses the exact same wiring setup as the `onenote` program.

Each time you press the controller, it will play the next chord in the programmed chord progression.

To build/flash the `chorder` program on Arduino Nano RP2040:

        tinygo flash -target nano-rp2040 ./chorder/

Launch one of the online synths. You should be able to use your new custom MIDI controller to make music.

Have fun!

### fourkeys

This MIDI controller sends four different notes.

- Connect a black cable from one pin of the green button on the breadboard to the breadboard's ground rail (-).

- Connect a colored cable from the other pin of the green button on the breadboard to pin D11 on the Arduino.

- Connect a black cable from one pin of the red button on the breadboard to the breadboard's ground rail (-).

- Connect a colored cable from the other pin of the red button on the breadboard to pin D10 on the Arduino.

- Connect a black cable from one pin of the yellow button on the breadboard to the breadboard's ground rail (-).

- Connect a colored cable from the other pin of the yellow button on the breadboard to pin D9 on the Arduino.

To build/flash the `fourkey` program on Arduino Nano RP2040:

        tinygo flash -target nano-rp2040 ./fourkey/

