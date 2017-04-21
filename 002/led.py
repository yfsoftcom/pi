#!/usr/bin/env python  
#encoding: utf-8  
import RPi.GPIO as GPIO
import time

# 针脚编号
channels = [18]   

def init():
  GPIO.setmode(GPIO.BOARD)  
  for x in channels:
    GPIO.setup(x, GPIO.OUT)

def on(i):  
  GPIO.output(channels[i], GPIO.HIGH)  

def off(i):  
  GPIO.output(channels[i], GPIO.LOW)  

# def ctrl(data):  
#   for i in channels:  
#     GPIO.output(i, data & 0x1)  
#     data = data >> 1  
#   pass  

def test(): 
  for i in xrange(10):  
    # ctrl(i) 
    if i%2 == 0:
      on(0)
    else:
      off(0)
    time.sleep(0.2)

def clean():  
  GPIO.cleanup()

