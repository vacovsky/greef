package monitor

import (
	"math"

	"github.com/davecgh/go-spew/spew"
)

func average(vals []float64) float64 {
	sum := 0.0
	for _, i := range vals {
		sum += i
	}
	return (sum / float64(len(vals)))
}

func removeFromSlice(s []float64, i int) []float64 {
	newSlice := []float64{}
	found := false

	spew.Dump(i, s, newSlice)

	for _, v := range s {
		if s[i] != v {
			newSlice = append(newSlice, v)
		} else if almostEqual(s[i], v) && !found {
			newSlice = append(newSlice, v)
			found = true
		}
	}
	return newSlice
}

func almostEqual(a, b float64) bool {
	float64EqualityThreshold := 0.001
	return math.Abs(a-b) <= float64EqualityThreshold
}
