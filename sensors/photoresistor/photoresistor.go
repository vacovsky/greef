package photoresistor

import "github.com/vacovsky/greef/sensors/mcp3008"

func Read(channel int, spiDev string, vref float64) float64 {
	volts, _ := mcp3008.ReadChannel(channel, spiDev, vref)
	return volts
}
