import RPi.GPIO as GPIO
import time
gpio_in1 = 20
gpio_in2 = 21
LOW = 0
HIGH = 1

pin_mtdriver = [in1, in2]
GPIO.setmode(GPIO.BCM)
GPIO.setup(pin_mtdriver, GPIO.OUT)

GPIO.output(gpio_in1, LOW)
GPIO.output(gpio_in2, LOW)
time.sleep(3)
GPIO.output(gpio_in1, LOW)
GPIO.output(gpio_in2, HIGH)
time.sleep(3)
GPIO.output(gpio_in1, HIGH)
GPIO.output(gpio_in2, LOW)
time.sleep(3)
GPIO.output(gpio_in1, HIGH)
GPIO.output(gpio_in2, HIGH)

GPIO.cleanup()