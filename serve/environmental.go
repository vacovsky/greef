package serve

import (
	"encoding/json"
	"net/http"

	"github.com/vacovsky/greef/data/queries"
)

func getData(w http.ResponseWriter, r *http.Request) {
	ep, err := queries.GetLastEnvironmentalParams()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getTemperatureChartData(w http.ResponseWriter, r *http.Request) {
	ep := queries.GetTemperaturesChartData()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getPHChartData(w http.ResponseWriter, r *http.Request) {
	ep := queries.GetPHChartData()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getTdsChartData(w http.ResponseWriter, r *http.Request) {
	ep := queries.GetTdsChartData()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getLightingChartData(w http.ResponseWriter, r *http.Request) {
	ep := queries.GetLightingChartData()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getCo2ChartData(w http.ResponseWriter, r *http.Request) {
	ep := queries.GetCo2ChartData()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
