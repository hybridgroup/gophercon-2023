# Gophercon.EU 2022

The is the repository for the hardware hack session at Gophercon.EU 2022.

https://gophercon.eu/

### Please return all equipment when you are finished for the next person. Thank you!

## Installation

### Clone this repo

First use git to clone this repo to your local machine:

```
git clone https://github.com/hybridgroup/gopherconeu-2022.git
cd gopherconeu-2022
```

### Install Go 1.18

If somehow you have not installed Go 1.18 on your computer already, you can download it here:

https://golang.org/dl/

Now you are ready to install TinyGo.

### Install TinyGo

You will need to install the TinyGo 0.25.0-beta1 in order to do today's activities.

You can find it at https://github.com/tinygo-org/tinygo/releases/tag/v0.25.0-beta1

Download the version for your OS, then follow these instructions here for your operating system:

<b>MacOS</b>

https://tinygo.org/getting-started/install/macos/#alternative-installation

Substitute the v0.25.0-beta1 in the above commands.

<b>Linux</b>

https://tinygo.org/getting-started/install/linux/#ubuntudebian

Substitute the v0.25.0-beta1 in the above commands.

<b>Windows</b>

https://tinygo.org/getting-started/install/windows/#manual-install

Substitute the v0.25.0-beta1 in the above commands.


## Activities

### TinyGo IoT sensors

We have some Arduino Nano RP2040 Connect IoT microcontroller boards for each person to use for the activity.

https://store.arduino.cc/collections/boards/products/arduino-nano-rp2040-connect

These can be programmed using TinyGo.

There are 24 Grove sensor kits that you can use for the activity.

Ready to try this out? Go to [./sensor/arduino/](./sensor/arduino/) to get started.

### Gopherdrone

We have three DJI Tello drones for your Go-powered flying activities:

In additional to these drones to be to be coded/flown, we also have Dualshock3-clone controllers for flight control. Post your awesome videos using hashtag #gopherconeu

Want to get airborne? Go to [./drone/tello/](./drone/tello/).

### Gopherbot

We have four Gopherbots for everyone to take turns playing with.

Gopherbot is a robotic gopher plushie that can be programmed using TinyGo.

Check out https://github.com/hybridgroup/gopherbot for more info.

### TinyGo Music Jam

Make your own electronic musical instruments using TinyGo and the Raspberry Pi Nano boards.

Thanks to the USB-MIDI support, you can turn your board into a tiny digital musical instrument controller.

Add some copper foil tape and high-ohm resistors when we say become part of the music we mean it!

We can make beautiful music together, just go to [./musicjam/](./musicjam/).

## License

Copyright (c) 2015-2022 The Hybrid Group and friends. Licensed under the MIT license.
