package ks0429

import "github.com/vacovsky/greef/sensors/mcp3008"

// ReadTdsPpm reads TDS of water
func ReadTdsPpm(channel int, spiDev string, vref, temperature float64) float64 {
	result := 0.0
	v, _ := mcp3008.ReadChannel(channel, spiDev, vref)
	compensationCoefficient := 1.0 + 0.02*(convertFtoC(temperature)-25.0)
	compensationVoltage := v / compensationCoefficient
	result = (133.42*compensationVoltage*compensationVoltage*compensationVoltage - 255.86*compensationVoltage*compensationVoltage + 857.39*compensationVoltage) * 0.5
	return result
}

func convertFtoC(f float64) float64 {
	return (f - 32) / 1.8
}
