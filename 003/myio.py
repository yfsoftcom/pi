load_gpio = False

try:
    import RPi.GPIO as GPIO
except:
    print('import gpio error~')
else:
    load_gpio = True

# ChannelIOs = [ 4, 5, 6, 12, 13, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 ] 
ChannelIOs = [  ] 

def is_setup(channel):
    return int(channel) in ChannelIOs

def clean():  
    del ChannelIOs[:]
    GPIO.cleanup()

def setup(channels):
    init()
    arr = channels.split(',')
    for x in arr:
        channel = int(x)
        if is_setup(channel):
            continue
        ChannelIOs.append(channel)
        GPIO.setup(channel, GPIO.OUT)

def init():
    GPIO.setmode(GPIO.BCM)
    GPIO.setwarnings(False)
    # for x in ChannelIOs:
    #     GPIO.setup(x, GPIO.OUT)

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
