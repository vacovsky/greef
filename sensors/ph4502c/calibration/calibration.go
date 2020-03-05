package main

import (
	"log"
	"time"

	"github.com/vacovsky/greef/sensors/mcp3008"
)

func main() {
	for {
		calibrate()
		time.Sleep(time.Millisecond * 250)
	}
}

var (
	spiDev      = "/dev/spidev0.0"
	testChannel = 7
	tChannel    = 1
	pChannel    = 0
	vref        = 5.0
	data        = []float64{}
)

func calibrate() {

	// ReadPH ADC channel and return pH detected by the probe
	pVolts, _ := mcp3008.ReadChannel(pChannel, spiDev, vref)

	data = append(data, pVolts)
	time.Sleep(time.Millisecond * 100)

	log.Println("Average:", average(), "\t\tpH volts:\t", pVolts)

}

func average() float64 {
	sum := 0.0
	for _, i := range data {
		sum += i
	}
	return (sum / float64(len(data)))
}
