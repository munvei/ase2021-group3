import sound
import RPi.GPIO as GPIO
import time
import curtain
import requests
pin = 23
gpio_sensor1 = 24
gpio_sensor2 = 25
pins = [pin, gpio_sensor1, gpio_sensor2]
GPIO.setmode(GPIO.BCM)
GPIO.setup(pins, GPIO.IN)
tmp = curtain.open()
dst = tmp + 30
try:
    while True:
        data1 = GPIO.input(gpio_sensor1)
        if(data1 == 1):
            curtain.stop()

        val = GPIO.input(pin)
        print("input:", val)
        time.sleep(0.25)
        cmptime = time.time()
        if(cmptime > dst):
            sound.buzzer()
            break
        if(val == 1):
            payload = {'msg': 'test messages'}
            requests.post("http://100.26.133.191/line", data=payload)
            break
                
except keyboardInterrupt:
    GPIO.output()
    GPIO.cleanup()
