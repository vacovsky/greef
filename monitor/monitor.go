package monitor

import (
	"log"
	"sync"
	"time"

	"github.com/vacovsky/greef/data"
	"github.com/vacovsky/greef/data/queries"
	"github.com/vacovsky/greef/oleddisplay"
)

var (
	currentEnvironment = data.EnvironmentalParameters{}
	mux                sync.Mutex
	checkInterval      = 5
	adcVRef            = 5.0
	adcSpiDev          = "/dev/spidev0.0"
)

func init() {
	mux = sync.Mutex{}
}

// Start the monitoring routines
func Start() {
	go func() {
		go monitorAmbientAir()
		go monitorWaterPH()
		go monitorAirCO2()
		go monitorAmbientLight()
		go monitorAmbientSumpLight()
		go monitorWaterTemperature()

		// TODO: TDS probe voltage interferes with ph probe.
		// The power for one must be disabled while other is active.
		// go monitorWaterTds()
	}()

	go func() {
		log.Println("Init database logging loop.", "Next write in 5 minutes.")
		time.Sleep(time.Minute * 5)
		for {
			mux.Lock()
			currentEnvironment.ReadingDateTime = time.Now()
			queries.SaveEnvironmentalParams(currentEnvironment)
			mux.Unlock()
			time.Sleep(time.Minute * 5)
		}
	}()

	// I2C OLED Display
	go oleddisplay.LaunchDisplay()
}
