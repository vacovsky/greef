package ph4502c

import (
	"log"
	"os"
	"strconv"

	"github.com/vacovsky/greef/sensors/mcp3008"
)

var (
	ph1, voltage1, ph2, voltage2 float64
)

func init() {
	var err error
	ph1, err = strconv.ParseFloat(os.Getenv("GREEF_PH1"), 64)
	voltage1, err = strconv.ParseFloat(os.Getenv("GREEF_PH1VOLTAGE"), 64)

	ph2, err = strconv.ParseFloat(os.Getenv("GREEF_PH2"), 64)
	voltage2, err = strconv.ParseFloat(os.Getenv("GREEF_PH2VOLTAGE"), 64)

	if err != nil {
		log.Fatal("Enable to load PH calibration values from environment variables.  See documentation to address this issue.")
	}
}

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

	slope, yIntercept := helpers.slopeIntercept(ph1, voltage1, ph2, voltage2)
	return (slope*v + yIntercept)

	// y=mx+b DO YOUR OWN CALIBRATION! return (-5.70*v + 21.34)
	// return (-5.972*v + 31.105)
}
