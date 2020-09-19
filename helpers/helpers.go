package helpers

import (
	"math"
	"sort"
)

func Mean(vals []float64) float64 {
	sum := 0.0
	for _, i := range vals {
		sum += i
	}
	return (sum / float64(len(vals)))
}

func Median(vals []float64) float64 {
	if len(vals) < 1 {
		return 0.0
	}
	sorted := vals
	sort.Float64s(sorted)

	if len(sorted) >= 2 {
		index := len(sorted) / 2

		if len(sorted)%2 == 0 {
			return (sorted[index-1] + sorted[index]) / float64(2.0)
		}
		return sorted[index]
	}
	return vals[0]
}

func RemoveFromSlice(s []float64, i int) []float64 {
	newSlice := []float64{}
	found := false

	for _, v := range s {
		if s[i] != v {
			newSlice = append(newSlice, v)
		} else if AlmostEqual(s[i], v) && !found {
			newSlice = append(newSlice, v)
			found = true
		}
	}
	return newSlice
}

func AlmostEqual(a, b float64) bool {
	float64EqualityThreshold := 0.001
	return math.Abs(a-b) <= float64EqualityThreshold
}

func SlopeIntercept(known1, voltage1, known2, voltage2 float64) (float64, float64) {

	slope, yIntercept := 0.0, 0.0

	// find the slope
	slope = (voltage2 - voltage1) / (known2 - known1)

	// find the y-intercept
	// y = mx+b
	yIntercept = voltage2 - (slope * known2)

	return slope, yIntercept

}
