package queries

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/vacovsky/greef/data"
)

var chartHours = 170

func init() {
	if os.Getenv("GREEF_CHART_HOURS") != "" {
		chartHours, _ = strconv.Atoi(os.Getenv("GREEF_CHART_HOURS"))
	}
}

func GetLastEnvironmentalParams() (data.EnvironmentalParameters, error) {
	var result data.EnvironmentalParameters
	var err error
	data.Service().Last(&result)
	return result, err
}

func GetEnvironmentalParams(limit int) ([]data.EnvironmentalParameters, error) {
	var result []data.EnvironmentalParameters
	var err error
	limit = 4032
	data.Service().Limit(limit).Find(&result)
	return result, err
}

func GetPHChartData() data.Chart {
	rawdata := []data.EnvironmentalParameters{}
	labels := []int64{}
	pHvals := []string{}

	data.Service().Where("reading_date_time > ?", time.Now().Add(-time.Duration(chartHours)*time.Hour)).Find(&rawdata).Order("time_stamp desc")

	for _, v := range rawdata {
		labels = append(labels, v.ReadingDateTime.Unix()*1000)
		pHvals = append(pHvals, fmt.Sprintf("%.2f", v.WaterPH))
	}
	chart := data.Chart{
		Series: []string{"Power of Hydrogen (pH)"},
		Labels: labels,
		Data:   [][]string{pHvals},
	}
	return chart
}

func GetTdsChartData() data.Chart {
	rawdata := []data.EnvironmentalParameters{}
	labels := []int64{}
	tdsVals := []string{}

	data.Service().Where("reading_date_time > ?", time.Now().Add(-time.Duration(chartHours)*time.Hour)).Find(&rawdata).Order("time_stamp desc")

	for _, v := range rawdata {
		labels = append(labels, v.ReadingDateTime.Unix()*1000)
		tdsVals = append(tdsVals, fmt.Sprintf("%.2f", v.WaterTdsPpm))
	}
	chart := data.Chart{
		Series: []string{"Total Dissolved Solids PPM (TDS)"},
		Labels: labels,
		Data:   [][]string{tdsVals},
	}
	return chart
}

func GetCo2ChartData() data.Chart {
	rawdata := []data.EnvironmentalParameters{}
	labels := []int64{}
	co2Vals := []string{}

	data.Service().Where("reading_date_time > ?", time.Now().Add(-time.Duration(chartHours)*time.Hour)).Find(&rawdata).Order("time_stamp desc")

	for _, v := range rawdata {
		labels = append(labels, v.ReadingDateTime.Unix()*1000)
		co2Vals = append(co2Vals, fmt.Sprintf("%.2f", v.AmbientAirCO2))
	}
	chart := data.Chart{
		Series: []string{"Air CO2 PPM"},
		Labels: labels,
		Data:   [][]string{co2Vals},
	}
	return chart
}

func GetLightingChartData() data.Chart {
	rawdata := []data.EnvironmentalParameters{}
	labels := []int64{}
	vals := []string{}

	data.Service().Where("reading_date_time > ?", time.Now().Add(-time.Duration(chartHours)*time.Hour)).Find(&rawdata).Order("time_stamp desc")

	for _, v := range rawdata {
		labels = append(labels, v.ReadingDateTime.Unix()*1000)
		vals = append(vals, fmt.Sprintf("%.0f", 100.0-((v.AmbientLight/5)*100)))
	}
	chart := data.Chart{
		Series: []string{"Ambient Light"},
		Labels: labels,
		Data:   [][]string{vals},
	}
	return chart
}

func GetTemperaturesChartData() data.Chart {
	rawdata := []data.EnvironmentalParameters{}
	labels := []int64{}

	ambientTempStr := []string{}
	waterTempStr := []string{}
	ambientHumStr := []string{}

	data.Service().Where("reading_date_time > ?", time.Now().Add(-time.Duration(chartHours)*time.Hour)).Find(&rawdata).Order("time_stamp desc")

	var curWT float64
	var curAT float64
	var curAH float64

	for _, v := range rawdata {
		labels = append(labels, v.ReadingDateTime.Unix()*1000)

		// remove dud readings from the chart to smooth it out
		if v.AmbientHumidity != 0 {
			curAH = v.AmbientHumidity
		}
		if v.AmbientAirTemperature != 32 {
			curAT = v.AmbientAirTemperature
		}

		curWT = v.WaterTemperature

		ambientTempStr = append(ambientTempStr, fmt.Sprintf("%.2f", curAT))
		waterTempStr = append(waterTempStr, fmt.Sprintf("%.2f", curWT))
		ambientHumStr = append(ambientHumStr, fmt.Sprintf("%.0f", curAH))

	}
	chart := data.Chart{
		Series: []string{"Water Temperature", "Ambient Air Temperature", "Ambient Humidity"},
		Labels: labels,
		Data:   [][]string{waterTempStr, ambientTempStr, ambientHumStr},
	}
	return chart
}

func SaveEnvironmentalParams(params data.EnvironmentalParameters) {
	data.Service().Save(&params)
}
