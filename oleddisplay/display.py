from luma.core.interface.serial import i2c
from luma.core.render import canvas
from luma.oled.device import ssd1306, ssd1325, ssd1331, sh1106
import time
import json
import apicli

serial = i2c(port=1, address=0x3C)
#  device = ssd1306(serial, rotate=0)
device = sh1106(serial, rotate=0)

displayData = None


while True:
# Box and text rendered in portrait mode
    with canvas(device) as draw:
        #  draw.rectangle(device.bounding_box, outline="white", fill="black")
        displayData = json.loads(apicli.read_latest_vals())
        if displayData != None:
            print(displayData)
            draw.text((5, 5), "W. T: " + str(round(displayData['WaterTemperature'], 2)) + " F" ,
                  fill='white')
            draw.text((5, 20), "W. pH: " + str(round(displayData['WaterPH'], 2)),
                  fill="white")
            draw.text((5, 35), "W. TDS: " + str(round(displayData['WaterTdsPpm'], 0)),
                  fill="white")
            draw.text((5, 50), "A. T/H: " + str(round(displayData['AmbientAirTemperature'], 1)) + " F / " + str(round(displayData['AmbientHumidity'], 1)) +" %",
                  fill="white")

    time.sleep(300)

