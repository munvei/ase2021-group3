import sound
import RPi.GPIO as GPIO
import time
import curtain
import requests
import sys

pin = 23
gpio_sensor1 = 24
gpio_sensor2 = 25
gpio_led = 26
pins = [pin, gpio_sensor1, gpio_sensor2]
GPIO.setmode(GPIO.BCM)
GPIO.setup(pins, GPIO.IN)
GPIO.setup(gpio_led, GPIO.OUT)
interval = 5

#アラーム鳴動と起床確認通知
def cnt(dst):
      flag = 0
      while True:
        cmptime = time.time()
        val = GPIO.input(pin)
        print("input:", val)
        time.sleep(0.25)
        GPIO.output(gpio_led, 1)

        if(cmptime > dst and flag == 0):
          flag = 1
          sound.buzzer()
#サーバ再起動時要確認
        if(val == 1):
          GPIO.output(gpio_led, 0)
          payload = {'msg': 'test messages'}
          requests.post("http://54.173.221.236/line", data=payload)
          return 0 

#カーテンの状態を検知するセンサの値を監視
def sensor(sensor_num):
  try:
      dst = time.time() + 3
      while True:
          #data = GPIO.input(gpio_sensor2)
          data = GPIO.input(sensor_num)
          cmptime = time.time()
          if(data == 1 or (cmptime > dst)):
              tmp = curtain.stop()
              dst = tmp + interval
              if(sys.argv [1] == "open"):
                cnt(dst)
              break
  except KeyboardInterrupt:
     GPIO.output()
     GPIO.cleanup()

#モータ動作命令 open:右回り close:左回り
if(sys.argv [1] == 'open' or sys.argv [1] == 'open_only'):
  curtain.open()
  sensor(gpio_sensor1)
elif(sys.argv [1] == 'close'):
  curtain.close()
  sensor(gpio_sensor2)

  
