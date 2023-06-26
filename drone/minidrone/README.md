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

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS3 controller to your computer. The controls are as follows:

* Triangle    - Takeoff
* X           - Land
* Left stick  - altitude
* Right stick - direction


IMPORTANT NOTE: you must press the "P3" button when your program first runs for the "clone" joysticks we are using to fully turn on.

**macOS**

`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

`go run ./step5/main.go [Bluetooth ID]`

**Linux**

`sudo apt-get install libsdl2-dev`

`go run step5/main.go [MAC address]`

**Windows**:

1. Install mingw-w64 from [Mingw-builds](https://github.com/niXman/mingw-builds-binaries/releases). A 7z archive extractor software might be needed which can be downloaded [here](https://www.7-zip.org/download.html). In this example, we extract the content, which is `mingw64`, into `C:\`.
2. Download and install `SDL2-devel-[version]-mingw.zip` files from https://github.com/libsdl-org/SDL/releases.
    * Extract the SDL2 folder from the archive using a tool like [7zip](http://7-zip.org)
    * Inside the extracted SDL2 folder, copy the `i686-w64-mingw32` and/or `x86_64-w64-mingw32` into mingw64 folder e.g. `C:\mingw64`
3. Setup `Path` environment variable
    * Put mingw-w64 binaries location into system `Path` environment variable (e.g. `C:\mingw64\bin`)
4. Close and open terminal again so the new `Path` environment variable takes effect. Now we should be able to run `go build` inside the project directory.
5. Download and install SDL2 runtime libraries from https://github.com/libsdl-org/SDL/releases. Extract and copy the `.dll` file into the project directory. After that, the program should become runnable.

`go run ./step5/main.go [MAC address]`

<hr>

### keyboard/main.go

Control the Minidrone with your keyboard!

- [, ] control take off and landing
- w, s, a, d control moving forward, backward, strafe left, and strafe right
- i, k, j, l control moving up, down, turning counter clockwise, and clockwise
- t, g, f, h control front flip, back flip, left flip, right flip
- r stop all movement on the tello to allow it to simply hover
