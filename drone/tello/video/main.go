package main

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	gobot "gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/dji/tello"
)

var drone = tello.NewDriver("8888")

func main() {
	var currentFlightData *tello.FlightData

	work := func() {
		mplayer := exec.Command("mplayer", "-fps", "60", "-")
		mplayerIn, _ := mplayer.StdinPipe()
		configureVideoEvents(mplayerIn)
		if err := mplayer.Start(); err != nil {
			fmt.Println(err)
			return
		}

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			fd := data.(*tello.FlightData)
			currentFlightData = fd
		})

		gobot.Every(1*time.Second, func() {
			printFlightData(currentFlightData)
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func configureVideoEvents(mplayerIn io.WriteCloser) {
	drone.On(tello.ConnectedEvent, func(data interface{}) {
		fmt.Println("Connected")
		drone.StartVideo()
		drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
		gobot.Every(100*time.Millisecond, func() {
			drone.StartVideo()
		})
	})

	drone.On(tello.VideoFrameEvent, func(data interface{}) {
		pkt := data.([]byte)
		if _, err := mplayerIn.Write(pkt); err != nil {
			fmt.Println(err)
		}
	})
}

func printFlightData(d *tello.FlightData) {
	if d.BatteryLow {
		fmt.Printf(" -- Battery low: %d%% --\n", d.BatteryPercentage)
	}

	displayData := `
Battery:		%d%%
Height:         %d
Ground Speed:   %d

`
	fmt.Printf(displayData, d.BatteryPercentage, d.Height, d.GroundSpeed)
}
