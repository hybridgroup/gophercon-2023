# Sphero Ollie

![Sphero Ollie](../../images/ollie.jpg)

The Sphero Ollie, Sphero SPRK+, and Sphero BB-8 all use the same API. However,
they have separate Gobot drivers to accommodate their other differences.

## What you need

    - Sphero Ollie
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux, macOS, or Windows

## Installation

Since this code uses the TinyGo Bluetooth package, you may have some specific installation requirements for your platform.

## Running the code

When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the Ollie using the Bluetooth interface.

On Linux and Windows you will use the MAC address of the device to connect.

On macOS you must use the Bluetooth ID of the device to connect.

Therefore, you must know the correct name and then MAC address or ID for that robot in order to connect to it.

The number of the robot should be listed on the side of it. You can lookup the correct name/MAC address in the following table.

|Number|Name|MAC Address|
|------|----|-----------|
|1|2B-C09F|D1:FE:13:8D:C0:9F|
|2|2B-6603|C2:FC:CE:65:66:03|
|3|2B-5F6A|D6:5F:69:D6:5F:6A|
|4|2B-DE50|E0:4A:58:0C:DE:50|

To find out the unique Bluetooth ID assigned to that device from macOS, you can use the Bluetooth scanner located in the `tools/blescanner` directory of this repo.

## Code

### step1

This tests that the Sphero Ollie is connected correctly to your computer, by blinking the built-in LED.

```
go run ./step1/ [MAC address or Bluetooth ID]
```

### step2

Rolls around at random.

```
go run ./step2/ [MAC address or Bluetooth ID]
```

### step3

Gets collision notifications from robot.

```
go run ./step3/ [MAC address or Bluetooth ID]
```

### step4

Control robot using keyboard.

```
go run ./step4/ [MAC address or Bluetooth ID]
```

`w` - forward
`a` - left
`s` - backward
`d` - right
`ESC`, `CTRL-C` - exit

### step5

This step has us receiving a heartbeat signal from the "base station" using the MQTT machine to machine messaging protocol. No additional hardware needs to be connected. 

You will need the server location of the MQTT server to use for the base station.

When the heartbeat data is received from the base station, the built-in LED will change color.


```
go run ./step5/ [MAC address or Bluetooth ID] [need this info]
```

### step6

Control robot using keyboard to collect data and send to base station.

```
go run ./step6/ [MAC address or Bluetooth ID] [need this info]
```

### step7

Control robot using keyboard to collect data and send to base station.

```
go run ./step7/ [MAC address or Bluetooth ID] [need this info]
```

## License

Copyright (c) 2015-2023 The Hybrid Group and friends. Licensed under the MIT license.
