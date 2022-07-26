package main

import (
	"image/color"
	"machine"
	"math/rand"
	"strconv"
	"time"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"

	"tinygo.org/x/drivers/net/mqtt"
	"tinygo.org/x/drivers/wifinina"
)

var (
	green   = machine.D12
	red  = machine.D10
	button = machine.D11
	touch  = machine.D9
	bzrPin = machine.D8

	bzr      buzzer.Device
	dial     = machine.ADC{machine.ADC0}
	pwm      = machine.PWM2 // PWM2 corresponds to Pin D10.
	redPwm uint8

	dialValue  uint16
	buttonPush bool
	touchPush  bool
)

var (
	// these are the default pins for the Arduino Nano33 IoT.
	spi = machine.NINA_SPI

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device
	topic   = "tinygo"

	mqttClient mqtt.Client
)

// access point info.
var (
	ssid string
	pass string
)

// IP address of the MQTT broker to use. Replace with your own info, if so desired.
var server = "tcp://test.mosquitto.org:1883"

func main() {
	initDevices()

	initAdaptor()
	connectToAP()

	connectToMQTT()

	go handleDisplay()
	go publishToMQTT()

	for {
		dialValue = dial.Get()
		pwm.Set(redPwm, uint32(dialValue))

		buttonPush = button.Get()
		if buttonPush {
			green.High()
		} else {
			green.Low()
		}

		touchPush = touch.Get()
		if touchPush {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 50)
	}
}

func initDevices() {
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
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
	redPwm, err = pwm.Channel(red)
	if err != nil {
		println("failed to configure PWM channel")
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

		time.Sleep(200 * time.Millisecond)
	}
}

func initAdaptor() {
	// Configure SPI for 8Mhz, Mode 0, MSB First
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	// Init esp32
	adaptor = wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()
}

// connect to access point
func connectToAP() {
	time.Sleep(2 * time.Second)
	println("Connecting to " + ssid)
	err := adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
	if err != nil { // error connecting to AP
		failMessage(err.Error())
	}

	println("Connected.")

	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		println(err.Error())
		time.Sleep(1 * time.Second)
	}
	println(ip.String())
}

func connectToMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server).SetClientID("tinygo-client-" + randomString(10))

	println("Connectng to MQTT...")
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}
}

func publishToMQTT() {
	for {
		println("Publishing MQTT message...")
		data := "{\"e\":[{ \"dv\":" +
			strconv.Itoa(int(dialValue)) +
			", \"bp\":" +
			strconv.FormatBool(buttonPush) +
			", \"tp\":" +
			strconv.FormatBool(touchPush) +
			" }]}"

		token := mqttClient.Publish(topic, 0, false, []byte(data))
		token.Wait()
		if token.Error() != nil {
			println(token.Error().Error())
		}
		time.Sleep(time.Second)
	}
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
