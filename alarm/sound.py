import RPi.GPIO as GPIO
import time
chan = 21
freq = 8000
def buzzer():
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(chan, GPIO.OUT)

    pwm = GPIO.PWM(chan, freq)

    pwm.start(50)
    time.sleep(3)
    pwm.stop()
    GPIO.cleanup()

if __name__ == '__main__':
    buzzer()

print('module name:{}'.format(__name__))
