package serve

//
// import (
// "fmt"
// "image"
// "log"
// "net/http"
// "os"
// "time"
//
// "gocv.io/x/gocv"
// )
//
// var (
// err     error
// webcam  *gocv.VideoCapture
// frameID int
// )
//
// func camLoader() {
// // open webcam
// if len(os.Args) < 2 {
// fmt.Println(">> device /dev/video0 (default)")
// webcam, err = gocv.VideoCaptureDevice(0)
// } else {
// fmt.Println(">> file/url :: " + os.Args[1])
// webcam, err = gocv.VideoCaptureFile(os.Args[1])
// }
// if err != nil {
// log.Println("Error opening video capture device.", "Retrying in 30 seconds!")
// time.Sleep(time.Second * 30)
// defer camLoader()
// return
// }
// defer webcam.Close()
//
// // start capturing
// go getframes()
// }
//
// func streamVideo(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
// data := ""
// for {
// mutex.Lock()
// data = "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"
// mutex.Unlock()
// time.Sleep(33 * time.Millisecond)
// w.Write([]byte(data))
// }
// }
//
// func getframes() {
// img := gocv.NewMat()
// defer img.Close()
// for {
// if ok := webcam.Read(&img); !ok {
// fmt.Printf("Device closed\n")
// return
// }
// if img.Empty() {
// continue
// }
// frameID++
// gocv.Resize(img, &img, image.Point{}, float64(1), float64(1), 0)
// frame, _ = gocv.IMEncode(".jpg", img)
//
// }
// }
