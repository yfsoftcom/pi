load_gpio = False

try:
    import RPi.GPIO as GPIO
except:
    print('import gpio error~')
else:
    load_gpio = True

# channels = [ 4, 5, 6, 12, 13, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 ] 
channels = [ 18 ] 

def is_setup(channel):
    return int(channel) in channels

def clean():  
  GPIO.cleanup()

def init():
    GPIO.setmode(GPIO.BOARD)  
    GPIO.setwarnings(False)
    for x in channels:
        GPIO.setup(x, GPIO.OUT)

def on(i):  
    print('on:', i)
    GPIO.output(int(i), GPIO.HIGH)  

def off(i):  
    print('off:', i)
    GPIO.output(int(i), GPIO.LOW)  

if load_gpio:
    init()
else:
    print('cant load gpio')
