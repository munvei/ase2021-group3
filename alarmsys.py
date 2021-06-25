import ledswitch
import sound
import RPi.GPIO as GPIO
import time
pin = 24
GPIO.setmode(GPIO.BCM)
GPIO.setup(pin, GPIO.IN)
tmp = ledswitch.blinkled()
dst = tmp + 10
try:
    while True:
        val = GPIO.input(pin)
        print("input:", val)
        time.sleep(0.25)
        cmptime = time.time()
        if(cmptime > dst):
            sound.buzzer()
            break
        if(val == 1):
            break
                
except keyboardInterrupt:
    GPIO.cleanup()
