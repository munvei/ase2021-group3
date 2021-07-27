import sound
import RPi.GPIO as GPIO
import time
import curtain
import requests
import sys

pin = 23
gpio_sensor1 = 24
gpio_sensor2 = 25
pins = [pin, gpio_sensor1, gpio_sensor2]
GPIO.setmode(GPIO.BCM)
GPIO.setup(pins, GPIO.IN)
interval = 30

def cnt(dst):
      while True:
        cmptime = time.time()
        val = GPIO.input(pin)
        print("input:", val)
        time.sleep(0.25)

        if(cmptime > dst):
          if(sys.argv [1] == 'open'):
            sound.buzzer()

        if(val == 1):
          payload = {'msg': 'test messages'}
          requests.post("http://54.173.221.236/line", data=payload)
          return 0 
     
def sensor(sensor_num):
  try:
      dst = time.time() + 10
      while True:
          #data = GPIO.input(gpio_sensor2)
          data = GPIO.input(sensor_num)
          cmptime = time.time()
          if(data == 1 or (cmptime > dst)):
              tmp = curtain.stop()
              dst = tmp + interval
              cnt(dst)
              break
  except KeyboardInterrupt:
     GPIO.output()
     GPIO.cleanup()
 
if(sys.argv [1] == 'open' or sys.argv [1] == 'open_only'):
  curtain.open()
  sensor(24)
elif(sys.argv [1] == 'close'):
  curtain.close()
  sensor(25)
