# DJI Tello Drone

![Tello](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a2/Ryze_Tello.jpg/320px-Ryze_Tello.jpg)

The DJI Tello from Ryze Robotics uses a WiFi interface with a UDP-based API.

You can use [Gobot](https://github.com/hybridgroup/gobot) to control the drone from your notebook computer.

## What you need

    - DJI Tello
    - Dualshock 3 gamepad, or compatible
    - Personal computer with Go 1.18+ installed
    - Works on Linux, macOS, or Windows

Do you have a Gopher Badge or GoBadge? You can also use it to control your drone using the Flightbadge firmware! Look in the tutorials directory in the repo for your specific badge.

## Installation

Change directories into this one where the needed Go modules files are located. 

Any other dependencies are listed in the tutorial under the step where they are needed.

## Connecting to Tello as WiFi access point

To communicate with your Tello, you need to connect to it using WiFi. The Tello will act as a WiFi access point, and you can just connect to it like you would any other.

Look on the inside of the drone where the battery goes for a small label that has the name of that Tello access point. It will be named something like "TELLO-C48BF0". Note that you must remove the battery to see the label.

Make sure you have a battery in the Tello, and turn it on by pushing the button located on the side of the drone.

Once the LED on the front of the Tello starts blinking amber rapidly, it is ready for you to connect to.

Use your computer's normal WiFi connection to connect to your Tello which should now appear in the list of available access points. There is no password.

Make sure you do not connect to someone else's Tello by mistake.

Note that you will not be able to connect to the Internet at the same time as you are connected to your Tello. This usually means a lot of connecting and disconnecting.

Now you are ready to try running some code.

## Running the code

When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the Tello using the WiFi interface.

Therefore, you must be connected to the Tello drone which acts as a WiFi access point before you will be able to run any of the code.

Once your program is running and connects to the drone, the LED on the front of the drone will blink more slowly. If the program you are running will cause it to take off, make sure you get out of the way! :)

## Code

### step01/main.go

Let's start with a simple takeoff, and then land. Make sure the drone is turned on and you're connected to its wifi access point, then run the code.

```go run step1/main.go```

<hr>

### step02/main.go

The drone will hover and return some flight data info. Run this code:

```go run step2/main.go```

<hr>

### step03/main.go

**NOTE: Ctrl-C will now land the drone if you get in trouble!**

The drone can move forward, backward, to the right, and the left, all while maintaining a steady altitude. Run the code. 

```go run step3/main.go```

<hr>

### step04/main.go

The drone can perform flips while flying. Run the code.

```go run step4/main.go```

<hr>

### step04a/main.go

This is the same functionality as step04, but instead of using Metal Gobot now we switch to using Standard Gobot. Notice the way that Gobot provides some functions like `Every()` and `After()` to help manage the various events. Run the code.

```go run step4a/main.go```

<hr>

### step05/main.go

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS4 or DS3 controller to your computer. The controls are as follows:

* Start    - Takeoff
* X           - Land
* Left stick  - altitude
* Right stick - direction


IMPORTANT NOTE: if using DS3 controller you must press the "P3" button when your program first runs for the "clone" joysticks we are using to fully turn on.

`go run step5/main.go`

<hr>

### step06/main.go

Now that you have mastered the flight controls, let's grab the drone video feed. You'll want to make sure that you have mplayer installed first. Upon running the code, you should see an mplayer window open with the camera feed.

**macOS**:
`brew install mplayer`

**Ubuntu Linux**:
`sudo apt-get install mplayer`

NOTE: you might have to open port 11111 for UDP on your machine like this:
`sudo ufw allow 11111/udp`

**Windows**:
Download from https://oss.netfarm.it/mplayer/

```go run step6/main.go```

<hr>

### keyboard/main.go

Control the tello with your keyboard!

- [, ] control take off and landing
- w, s, a, d control moving forward, backward, strafe left, and strafe right
- i, k, j, l control moving up, down, turning counter clockwise, and clockwise
- t, g, f, h control front flip, back flip, left flip, right flip
- r stop all movement on the tello to allow it to simply hover
