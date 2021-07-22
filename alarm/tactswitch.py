import RPi.GPIO as GPIO
import time

pin = 23
def push():
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(pin, GPIO.IN)
  
    try:
        while True:
            val = GPIO.input(pin)
            print("input:", val)
            time.sleep(0.25)
    except keyboardInterrupt:
        GPIO.cleanup()
if __name__ == '__main__':
    push()
print('module name:{}'.format(__name__))
