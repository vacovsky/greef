package monitor

import (
	"fmt"
	"time"

	"github.com/vacovsky/greef/sensors/photoresistor"
)

func monitorAmbientSumpLight() {
	for {
		time.Sleep(time.Minute * 5)
	}
}

func monitorAmbientLight() {
	lastVals := []float64{}
	lastValsMaxLen := 600
	lastLight := 0.0
	lastIndex := 0

	for {
		lastLight = photoresistor.Read(4, adcSpiDev, adcVRef)
		if len(lastVals) < lastValsMaxLen {
			lastVals = append(lastVals, lastLight)
		} else {
			lastVals[lastIndex] = lastLight
			lastIndex++
			if lastIndex >= lastValsMaxLen {
				lastIndex = 0
			}
		}
		avg := helpers.mean(lastVals)
		mux.Lock()
		currentEnvironment.AmbientLight = avg
		mux.Unlock()
		fmt.Println("Light (Voltage):\t\t\t", lastLight)
		time.Sleep(time.Second * 5)
	}
}
