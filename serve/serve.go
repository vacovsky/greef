package serve

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var buffer = make(map[int][]byte)
var frame []byte
var mutex = &sync.Mutex{}

// Serve serves the server serve
func Serve() {
	host := "0.0.0.0:3000"
	// if len(os.Args) < 2 {
	// fmt.Println(">> device /dev/video0 (default)")
	// webcam, err = gocv.VideoCaptureDevice(0)
	// } else {
	// fmt.Println(">> file/url :: " + os.Args[1])
	// webcam, err = gocv.VideoCaptureFile(os.Args[1])
	// }
	// if err != nil {
	// time.Sleep(time.Second * 30)
	// webcam, err = gocv.VideoCaptureDevice(0)
	// fmt.Printf("Error opening capture device: \n")
	// }
	// defer webcam.Close()
	// start capturing
	// go getframes()

	// FIXME: this doesn't work and I'm not sure why...  I put the old stuff back above it.
	// go camLoader()
	fmt.Println("Capturing. Open http://" + host)

	// start http server
	// http.HandleFunc("/video", streamVideo)
	http.HandleFunc("/data", getData)
	http.HandleFunc("/tempchartdata", getTemperatureChartData)
	http.HandleFunc("/phchartdata", getPHChartData)
	http.HandleFunc("/tdschartdata", getTdsChartData)
	http.HandleFunc("/co2chartdata", getCo2ChartData)
	http.HandleFunc("/lightingchartdata", getLightingChartData)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/"+r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(host, nil))
}
