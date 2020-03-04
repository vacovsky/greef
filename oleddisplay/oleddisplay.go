package oleddisplay

import (
	"log"
	"os/exec"
	"time"
)

/*
http://codelectron.com/setup-oled-display-raspberry-pi-python/
$ sudo apt-get install python-dev python-pip libfreetype6-dev libjpeg-dev build-essential
$ sudo -H pip install --upgrade luma.oled

*/
func LaunchDisplay() {
	cmd := exec.Command("sudo", "python", "oleddisplay/display.py")
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		log.Println("Relaunching display in 30 seconds.")
		time.Sleep(time.Second * 30)
		LaunchDisplay()
	}
}
