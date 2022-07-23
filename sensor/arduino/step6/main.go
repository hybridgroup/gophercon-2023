package main

import (
	"image/color"
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/ssd1306"
)

var (
	blue = machine.D12
	green = machine.D10
	button = machine.D11
	touch = machine.D9
	bzrPin = machine.D8

	bzr buzzer.Device
	dial = machine.ADC{machine.ADC0}
	pwm = machine.PWM2 // PWM2 corresponds to Pin D10.
	greenPwm uint8

	dialValue  uint16
	buttonPush bool
	touchPush  bool
)

func main() {
	initDevices()

	go handleDisplay()

	for {
		dialValue = dial.Get()
		pwm.Set(greenPwm, uint32(dialValue))

		buttonPush = button.Get()
		if buttonPush {
			blue.High()
		} else {
			blue.Low()
		}

		touchPush = touch.Get()
		if touchPush {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func initDevices() {
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	touch.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	err := pwm.Configure(machine.PWMConfig{
		Period: 16384e3, // 16.384ms
	})
	if err != nil {
		println("failed to configure PWM")
		return
	}

	machine.InitADC()
	dial.Configure(machine.ADCConfig{})

	bzr = buzzer.New(bzrPin)
}

func handleDisplay() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  32,
	})

	display.ClearDisplay()

	black := color.RGBA{1, 1, 1, 255}

	for {
		display.ClearBuffer()

		val := strconv.Itoa(int(dialValue))
		msg := "dial: " + val
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 10, 20, msg, black)

		var radius int16 = 4
		if buttonPush {
			tinydraw.FilledCircle(&display, 16+32*0, 32-radius-1, radius, black)
		} else {
			tinydraw.Circle(&display, 16+32*0, 32-radius-1, radius, black)
		}
		if touchPush {
			tinydraw.FilledCircle(&display, 16+32*1, 32-radius-1, radius, black)
		} else {
			tinydraw.Circle(&display, 16+32*1, 32-radius-1, radius, black)
		}

		display.Display()

		time.Sleep(100 * time.Millisecond)
	}
}
