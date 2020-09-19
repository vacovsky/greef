package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/vacovsky/greef/helpers"
	"github.com/vacovsky/greef/sensors/ds18b20"
	"github.com/vacovsky/greef/sensors/ks0429"
	"github.com/vacovsky/greef/sensors/ph4502c"
)

func monitorWaterTemperature() {
	log.Println("Init monitorWaterTemperature")
	for {
		wt := ds18b20.ReadFahrenheit()
		mux.Lock()
		currentEnvironment.WaterTemperature = wt
		mux.Unlock()
		log.Println("Water Temp:\t\t", wt)
		time.Sleep(time.Second * 15)
	}
}

func monitorWaterPH() {
	lastVals := []float64{}
	lastValsMaxLen := 500
	lastph := 0.0
	lastIndex := 0

	for {
		lastph = ph4502c.ReadPH(0, adcSpiDev, adcVRef)
		if len(lastVals) < lastValsMaxLen {
			lastVals = append(lastVals, lastph)
		} else {
			lastVals[lastIndex] = lastph
			lastIndex++
			if lastIndex >= lastValsMaxLen {
				lastIndex = 0
			}
		}
		avg := helpers.Median(lastVals)
		mux.Lock()
		currentEnvironment.WaterPH = avg
		mux.Unlock()
		fmt.Println("pH:\t\t\t", lastph)
		time.Sleep(time.Millisecond * 500)
	}
}

func monitorWaterTds() {

	lastVals := []float64{}
	lastValsMaxLen := 100
	lastTds := 0.0
	lastIndex := 0

	for {
		lastTds = ks0429.ReadTdsPpm(2, adcSpiDev, adcVRef, currentEnvironment.WaterTemperature)
		if len(lastVals) < lastValsMaxLen {
			lastVals = append(lastVals, lastTds)
		} else {
			lastVals[lastIndex] = lastTds
			lastIndex++
			if lastIndex >= lastValsMaxLen {
				lastIndex = 0
			}
		}
		avg := helpers.Median(lastVals)
		mux.Lock()
		currentEnvironment.WaterTdsPpm = avg
		mux.Unlock()
		fmt.Println("TDS:\t\t\t", lastTds)
		time.Sleep(time.Second * 6)
	}
}
