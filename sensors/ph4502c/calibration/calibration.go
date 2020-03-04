package main

import (
	"log"
	"time"

	"github.com/vacovsky/greef/sensors/mcp3008"
)

func main() {
	for {
		calibrate()
		time.Sleep(time.Second * 5)
	}
}

var (
	spiDev      = "/dev/spidev0.0"
	testChannel = 7
	tChannel    = 1
	pChannel    = 0
	vref        = 5.0
)

func calibrate() {

	// ReadPH ADC channel and return pH detected by the probe
	pVolts, _ := mcp3008.ReadChannel(pChannel, spiDev, vref)

	time.Sleep(time.Second * 1)
	// ReadWaterTempF ADC channel and return temperature detected by the probe
	tVolts, _ := mcp3008.ReadChannel(tChannel, spiDev, vref)
	time.Sleep(time.Second * 1)
	testVolts, _ := mcp3008.ReadChannel(testChannel, spiDev, vref)

	log.Println("Test volts:\t", testVolts, "\t\tpH volts:\t", pVolts, " ----- ", "\t\tTemp volts:\t", tVolts)

}
