package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
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
)

func main() {
	initDevices()

	for {
		pwm.Set(greenPwm, uint32(dial.Get()))

		if button.Get() {
			blue.High()
		} else {
			blue.Low()
		}

		if touch.Get() {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 100)
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
	greenPwm, err = pwm.Channel(green)
	if err != nil {
		println("failed to configure PWM channel")
		return
	}

	machine.InitADC()
	dial.Configure(machine.ADCConfig{})

	bzr = buzzer.New(bzrPin)
}