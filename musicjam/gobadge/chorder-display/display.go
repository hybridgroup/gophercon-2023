package main

import (
	"machine"

	"image/color"
	"time"

	"tinygo.org/x/drivers/st7735"
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
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	display.FillScreen(black)
	msg := ""

	for {
		if msg != chords[index].name {
			tinyfont.WriteLine(&display, &freemono.Bold12pt7b, startx, starty, msg, black)
			msg = chords[index].name
			tinyfont.WriteLine(&display, &freemono.Bold12pt7b, startx, starty, msg, white)
		}

		tinydraw.FilledCircle(&display, startx+radius+48*0, 48-radius-1, radius, black)
		if pressed {
			tinydraw.FilledCircle(&display, startx+radius+48*0, 48-radius-1, radius, white)
		} else {
			tinydraw.Circle(&display, startx+radius+48*0, 48-radius-1, radius, white)
		}

		display.Display()

		time.Sleep(200 * time.Millisecond)
	}
}
