import RPi.GPIO as GPIO
import time
def blinkled():
    crtime = time.time()
    gpio_led = 17

    GPIO.setmode(GPIO.BCM)

    GPIO.setup(gpio_led, GPIO.OUT)

    GPIO.output(gpio_led, 1)
    time.sleep(5)
    GPIO.output(gpio_led, 0)

    GPIO.cleanup(gpio_led)
    return crtime

if __name__ == '__main__':
    blinkled()

print('module name:{}'.format(__name__))
