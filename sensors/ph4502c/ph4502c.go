package ph4502c

import (
	"github.com/vacovsky/greef/sensors/mcp3008"
)

// ReadPH ADC channel and return pH detected by the probe
func ReadPH(channel int, spiDev string, vref float64) float64 {
	volts, _ := mcp3008.ReadChannel(channel, spiDev, vref)
	return parsePH(volts)
}

// ReadWaterTempF ADC channel and return temperature detected by the probe
func ReadWaterTempF(channel int, spiDev string, vref float64) float64 {
	volts, _ := mcp3008.ReadChannel(channel, spiDev, vref)
	return parseTemp(volts)
}

//  y= -5.70 * x + 21.34

func parseTemp(t float64) float64 {
	return (t * 1.8) + 32
}

func parsePH(v float64) float64 {
	// y=mx+b DO YOUR OWN CALIBRATION! return (-5.70*v + 21.34)
	return (-6.2*v + 29.8)
}
