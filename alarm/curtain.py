import RPi.GPIO as GPIO
import time
gpio_in1 = 20
gpio_in2 = 21
LOW = 0
HIGH = 1

pin_mtdriver = [gpio_in1, gpio_in2]
GPIO.setmode(GPIO.BCM)
GPIO.setup(pin_mtdriver, GPIO.OUT)

def open():
    #crtime = time.time()
    GPIO.output(gpio_in1, HIGH)
    GPIO.output(gpio_in2, LOW)
    #return crtime

def stop():
    crtime = time.time()
    GPIO.output(gpio_in1, HIGH)
    GPIO.output(gpio_in2, HIGH)
    return crtime

def close():
    #crtime = time.time()
    GPIO.output(gpio_in1, LOW)
    GPIO.output(gpio_in2, HIGH)
    #return crtime

if __name__ == '__main__':
    open()
    close()

print('module name:{}'.format(__name__))
