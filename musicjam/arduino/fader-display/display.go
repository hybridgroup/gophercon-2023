package main

import (
	"machine"
	"strconv"

	"image/color"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

const (
	startx       = 10
	starty       = 20
	radius int16 = 6
)

func handleDisplay() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  64,
	})

	display.ClearDisplay()

	black := color.RGBA{1, 1, 1, 255}

	for {
		display.ClearBuffer()

		msg := strconv.Itoa(int(svol))
		tinyfont.WriteLine(&display, &freemono.Bold12pt7b, startx, starty, msg, black)

		if pressed {
			tinydraw.FilledCircle(&display, startx+radius+48*0, 48-radius-1, radius, black)
		} else {
			tinydraw.Circle(&display, startx+radius+48*0, 48-radius-1, radius, black)
		}

		display.Display()

		time.Sleep(200 * time.Millisecond)
	}
}
