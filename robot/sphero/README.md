# Sphero Ollie

The Sphero Ollie, Sphero SPRK+, and Sphero BB-8 all use the same API. However,
they have separate Gobot drivers to accommodate their other differences.

## What you need

    - Sphero Ollie
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux, macOS, or Windows

## Installation

Since this code uses the TinyGo Bluetooth package, you may have some specific installation requirements for your platform.

More info here...

## Running the code

When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the robot using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the MAC address of your Ollie as needed.

More info needed here...

## Code

### step1

This tests that the Sphero Ollie is connected correctly to your computer, by blinking the built-in LED.

```
go run ./step1/ AA:BB:CC:DD:EE
```

### step2

Rolls around at random.

```
go run ./step2/ AA:BB:CC:DD:EE
```

### step3

Gets collision notifications from robot.

```
go run ./step3/ AA:BB:CC:DD:EE
```

### step4

Control robot using keyboard arrow keys.

```
go run ./step4/ AA:BB:CC:DD:EE
```

### step5

This step has us receiving a heartbeat signal from the "base station" using the MQTT machine to machine messaging protocol. No additional hardware needs to be connected. 

You will need the server location of the MQTT server to use for the base station.

When the heartbeat data is received from the base station, the built-in LED will change color.


```
go run ./step5/ AA:BB:CC:DD:EE [need this info]
```

### step6

Control robot using keyboard to collect data and send to base station.

```
go run ./step6/ AA:BB:CC:DD:EE [need this info]
```

### step7

Control robot using keyboard to collect data and send to base station.

```
go run ./step7/ AA:BB:CC:DD:EE [need this info]
```

## License

Copyright (c) 2015-2022 The Hybrid Group and friends. Licensed under the MIT license.
