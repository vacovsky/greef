package mcp3008

// SPI works on Pi in 4.9.41-v7+ but breaks in 4.9.59-v7+
// sudo rpi-update b9becbbf3f48e39f719ca6785d23c53ee0cdbe49

import (
	"fmt"
	"log"

	"github.com/advancedclimatesystems/io/spi/microchip"
	"github.com/vacovsky/embd"
	"github.com/vacovsky/embd/convertors/mcp3008"
	gbspi "gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
	"golang.org/x/exp/io/spi"
)

func Init() {

}

// ACSReadChannel returns the channel reading from the MCP3008 chipi
func ACSReadChannel(channel int, spidev string, vref float64) (float64, int) {
	conn, err := spi.Open(&spi.Devfs{
		Dev:      spidev,
		Mode:     spi.Mode0,
		MaxSpeed: 1000000,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to open SPI device: %s", err))
	}
	defer conn.Close()

	adc := microchip.MCP3008{
		Conn: conn,
		Vref: vref,
	}
	v, err := adc.Voltage(channel)
	if err != nil {
		log.Println(err)
	}

	c, err := adc.OutputCode(channel)
	if err != nil {
		log.Println(err)
	}

	voltage := float64(v)
	return voltage, c
}

// EmbdReadChannel ses emdb library to read the SPI device.  This seems more consistent across devices.
func EmbdReadChannel(channel int, spidev string, vref float64) (float64, int) {
	var channelByte = byte(channel)
	const (
		speed = 1000000 //5000
		bpw   = 8
		delay = 0
	)
	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()
	spiBus := embd.NewSPIBus(embd.SPIMode0, channelByte, speed, bpw, delay)
	defer spiBus.Close()

	adc := mcp3008.New(mcp3008.SingleMode, spiBus)

	val, err := adc.AnalogValueAt(channel)
	if err != nil {
		log.Println("Unable to read from SPIDEV at channel ", channel, err)
	}
	voltage := float64(vref) * (float64(val) / 1023.0)
	return voltage, val
}

// GobotReadChannel ses emdb library to read the SPI device.  This seems more consistent across devices.
func GobotReadChannel(channel int, spidev string, vref float64) (float64, int) {
	v, n := 0.0, 0
	r := raspi.NewAdaptor()

	d := gbspi.NewMCP3008Driver(r)

	err := d.Start()
	if err != nil {
		fmt.Println(err)
	}
	defer d.Halt()
	n, err = d.Read(channel)
	if err != nil {
		fmt.Println("Unable to read from SPIDEV at channel ", channel)
		log.Println(err)
	}

	v = (float64(n) / 1023.0)

	return v, n
}

func ReadChannel(channel int, spidev string, vref float64) (float64, int) {
	// v, n := GobotReadChannel(channel, spidev, vref)
	// v, n := EmbdReadChannel(channel, spidev, vref)
	v, n := ACSReadChannel(channel, spidev, vref)

	//log.Println(channel, v, n)
	return v, n
}
