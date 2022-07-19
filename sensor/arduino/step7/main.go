package main

import (
	"machine"
	"math/rand"
	"time"

	// comes from "github.com/conejoninja/tinyfont/freemono"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/espat"
	"tinygo.org/x/drivers/espat/mqtt"
)

var (
	dialValue  uint16
	buttonPush bool
	touchPush  bool

	uart = machine.UART1
	tx   = machine.PA22
	rx   = machine.PA23

	adaptor *espat.Device
	topic   = "tinygo"
)

// access point info. Change this to match your WiFi connection information.
const ssid = "GDG 2019"
const pass = "gdgsummit"

// IP address of the MQTT broker to use. Replace with your own info, if so desired.
const server = "tcp://test.mosquitto.org:1883"

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	rand.Seed(time.Now().UnixNano())

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	machine.InitADC()
	machine.InitPWM()

	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	green := machine.PWM{machine.D10}
	green.Configure()

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	touch := machine.D9
	touch.Configure(machine.PinConfig{Mode: machine.PinInput})

	bzrPin := machine.D8
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	dial := machine.ADC{machine.A0}
	dial.Configure()

	// Init esp8266/esp32
	adaptor = espat.New(uart)
	adaptor.Configure()

	// first check if connected
	if connectToESP() {
		blue.High()
		println("Connected to wifi adaptor.")
		adaptor.Echo(false)

		blue.Low()
		connectToAP()
		blue.High()
	} else {
		println("")
		failMessage("Unable to connect to wifi adaptor.")
		return
	}

	opts := mqtt.NewClientOptions(adaptor)
	opts.AddBroker(server).SetClientID("tinygo-client-" + randomString(10))

	blue.Low()
	println("Connectng to MQTT...")
	cl := mqtt.NewClient(opts)
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}

	for {
		dialValue = dial.Get()
		green.Set(dialValue)

		buttonPush = button.Get()
		if !buttonPush {
			blue.Low()
		} else {
			blue.High()
			println("Publishing MQTT message...")
			data := []byte("{\"e\":[{ \"n\":\"hello\", \"sv\":\"world\" }]}")
			token := cl.Publish(topic, 0, false, data)
			token.Wait()
			if token.Error() != nil {
				println(token.Error().Error())
			}
		}

		touchPush = touch.Get()
		if touchPush {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 100)
	}

	// Right now this code is only reached when there is an error. Need a way to trigger clean exit.
	println("Disconnecting MQTT...")
	cl.Disconnect(100)

	println("Done.")
}

// connect to ESP8266/ESP32
func connectToESP() bool {
	for i := 0; i < 5; i++ {
		println("Connecting to wifi adaptor...")
		if adaptor.Connected() {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// connect to access point
func connectToAP() {
	println("Connecting to wifi network...")

	adaptor.SetWifiMode(espat.WifiModeClient)
	adaptor.ConnectToAP(ssid, pass, 10)

	println("Connected.")
	println(adaptor.GetClientIP())
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}
