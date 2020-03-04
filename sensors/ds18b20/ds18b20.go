package ds18b20

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read the value from the temperature probe and return the value in C
// /sys/devices/w1_bus_master1/28-03189779c5a9
func Read() float64 {
	var err error
	var b []byte
	fileName := os.Getenv("W1_SLAVE")
	if fileName == "" {
		fileName = "/sys/devices/w1_bus_master1/28-03189779c5a9/w1_slave"
	}
	var result float64

	if fileExists(fileName) {
		b, err = ioutil.ReadFile(fileName)
	} else {
		log.Println("Unable to find", fileName, "Cannot read Water Temp...")
	}
	if err != nil {
		log.Print(err)
	}
	if strings.Contains(string(b), "t=") {
		tempRaw := strings.Trim(
			strings.Split(string(b), "t=")[1], "\n ")
		result, err = strconv.ParseFloat(tempRaw, 64)
	} else {
		log.Println("Unable to read", fileName)
		result = 0
	}
	return result / 1000
}

// ReadFahrenheit returns temp sensor reading in F instead of C
func ReadFahrenheit() float64 {
	t := Read()
	log.Println(t)
	return (t * 1.8) + 32
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
