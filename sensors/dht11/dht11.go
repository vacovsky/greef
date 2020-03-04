package dht11

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// GetAmbientInfo returns ambient humidity and temperature if no errors occur
func GetAmbientInfo(dhtPin int) (float64, float64) {
	cmd := exec.Command("sudo", "./bin/dht11", strconv.Itoa(dhtPin))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	// spew.Dump(out.String())
	outStr := strings.Replace(out.String(), "\n", "", -1)
	temps := strings.Split(outStr, " ")
	ah, err := strconv.ParseFloat(temps[0], 64)
	at, err := strconv.ParseFloat(temps[1], 64)

	if err != nil {
		log.Println("error reading sensor: ", err)
		return 0, 0
	}
	return ah, at
}

// GetAmbientInfoF returns humidity and temp (in farenheit)
func GetAmbientInfoF(dhtPin int) (float64, float64) {
	h, t := GetAmbientInfo(dhtPin)
	return h, farenheit(t)
}

// ReadFahrenheit returns temp sensor reading in F instead of C
func farenheit(val float64) float64 {
	return (val * 1.8) + 32
}
