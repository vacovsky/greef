package monitor

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vacovsky/greef/sensors/dht11"
)

func monitorAmbientAir() {
	var airPin int
	var err error
	for {
		if os.Getenv("GREEF_DHT_PIN") == "" {
			airPin = 18
		} else {
			airPin, err = strconv.Atoi(os.Getenv("GREEF_DHT_PIN"))
		}
		if err != nil {
			log.Println(err)
		}
		hum, temp := dht11.GetAmbientInfoF(airPin)

		if almostEqual(hum, 0.0) && almostEqual(temp, 32.0) {
			hum, temp = currentEnvironment.AmbientHumidity, currentEnvironment.AmbientAirTemperature
		}
		mux.Lock()
		currentEnvironment.AmbientHumidity, currentEnvironment.AmbientAirTemperature = hum, temp
		mux.Unlock()
		time.Sleep(time.Minute * 5)
	}
}

func monitorAirCO2() {
	for {
		time.Sleep(time.Minute * 5)
	}
}
