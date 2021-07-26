import RPi.GPIO as GPIO
import time
chan = 22
def buzzer():
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(chan, GPIO.OUT)

    GPIO.output(chan, 1)
    time.sleep(3)
    GPIO.output(chan, 0)
    GPIO.cleanup()

if __name__ == '__main__':
    buzzer()

print('module name:{}'.format(__name__))
