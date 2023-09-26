# Parrot Minidrone

![Parrot Minidrone](https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/Rolling_Spider.jpg/320px-Rolling_Spider.jpg)

The Minidrone from Parrot uses a Bluetooth interface and programming API.

You can use [Gobot](https://github.com/hybridgroup/gobot) to control the drone from your notebook computer.

## What you need

    - Parrot Minidrone
    - Personal computer with Go 1.19+ installed
    - Works on Linux, macOS, or Windows

Do you have a Gopher Badge or GoBadge? You can also use it to control your drone using the Flightbadge firmware! Look in the tutorials directory in the repo for your specific badge.

## Installation

Change directories into this one where the needed Go modules files are located. 

Any other dependencies are listed in the tutorial under the step where they are needed.

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the Minidrone using the Bluetooth interface.

On Linux and Windows you will use the MAC address of the device to connect.

On macOS you must use the Bluetooth ID of the device to connect.

Therefore, you must know the correct name and then MAC address or ID for that robot in order to connect to it.

The name of the drone should be listed on the side of it. You can lookup the correct MAC address in the following table.

|Name|MAC Address|
|----|-----------|
|Travis_056919|E0:14:5D:F0:3D:A4|
|Drone 2|XX|
|Drone 3|XX|

To find out the unique Bluetooth ID assigned to that device from macOS, you can use the Bluetooth scanner located in the tools directory of this repo.

## Code

### step01/main.go

Let's start with a simple takeoff, and then land. Make sure the drone is turned on and you know the correct MAC address or name, then run the code.

```go run ./step1/main.go [MAC address or Bluetooth ID]```

<hr>

### step02/main.go

The drone will hover and return some flight data info. Run this code:

```go run ./step2/main.go [MAC address or Bluetooth ID]```

<hr>

### step03/main.go

**NOTE: Ctrl-C will now land the drone if you get in trouble!**

The drone can move forward, backward, to the right, and the left, all while maintaining a steady altitude. Run the code. 

```go run ./step3/main.go [MAC address or Bluetooth ID]```

<hr>

### step04/main.go

The drone can perform flips while flying. Run the code.

```go run ./step4/main.go [MAC address or Bluetooth ID]```

<hr>

### step04a/main.go

This is the same functionality as step04, but instead of using Metal Gobot now we switch to using Standard Gobot. Notice the way that Gobot provides some functions like `Every()` and `After()` to help manage the various events. Run the code.

```go run ./step4a/main.go [MAC address or Bluetooth ID]```

<hr>

### step05/main.go

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS4 or DS3 controller to your computer. The controls are as follows:

* Triangle    - Takeoff
* X           - Land
* Left stick  - altitude
* Right stick - direction


IMPORTANT NOTE: you must press the "P3" button when your program first runs for the "clone" DS3 joysticks we are using to fully turn on.

**macOS**

`go run ./step5/main.go [Bluetooth ID]`

**Linux**

`go run step5/main.go [MAC address]`

**Windows**:

`go run ./step5/main.go [MAC address]`

<hr>

### keyboard/main.go

Control the Minidrone with your keyboard!

- [, ] control take off and landing
- w, s, a, d control moving forward, backward, strafe left, and strafe right
- i, k, j, l control moving up, down, turning counter clockwise, and clockwise
- t, g, f, h control front flip, back flip, left flip, right flip
- r stop all movement on the tello to allow it to simply hover
