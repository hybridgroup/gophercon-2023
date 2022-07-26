# DJI Tello Drone

The DJI Tello from Ryze Robotics uses a WiFi interface with a UDP-based API.

You can use [Gobot](https://github.com/hybridgroup/gobot) to control the drone from your notebook computer.

## What you need

    - DJI Tello
    - Dualshock 3 gamepad, or compatible
    - Personal computer with Go 1.18+ installed
    - Works on Linux, macOS, or Windows

## Installation


## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the Tello using the WiFi interface.

Therefore, you must connect to the Tello drone which acts as a WiFi access point before you will be able to run any of the code.

Further instructions here...

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

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS3 controller to your computer. The controls are as follows:

* Triangle    - Takeoff
* X           -  Land
* Left stick  - altitude
* Right stick - direction

**macOS**
`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

`go run step5/main.go`

**Linux**
`sudo apt-get install libsdl2-dev`

**Windows**:

1. Download and install the mingw-w64 compiler:

https://sourceforge.net/projects/mingw-w64/files/mingw-w64/mingw-w64-release/mingw-w64-v8.0.2.tar.bz2/download

2. Download and install SDL2:

http://libsdl.org/release/SDL2-2.0.22-win32-x64.zip

- Extract the SDL2 folder from the archive using a tool like 7zip
- Inside the folder, copy the x86_64-w64-mingw32 files into your mingw-w64 folder e.g. C:\Program Files\mingw-w64\x86_64-8.0.2-win32-seh-rt_v6\mingw64

3. Setup Path environment variable:

Put your mingw-w64 binaries location into your system Path environment variable. e.g C:\Program Files\mingw-w64\x86_64-8.0.2-win32-seh-rt_v6\mingw64\x86_64-w64-mingw32\bin

4. Install go-sdl2:

```
go get -v github.com/veandco/go-sdl2/sdl@master
```

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
