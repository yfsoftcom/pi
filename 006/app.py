# -*- coding: utf-8 -*-  
import paho.mqtt.client as mqtt
from myio import load_gpio, on, off, is_setup, clean, setup
import json

# 连接成功回调函数
def on_connect(client, userdata, flags, rc):
    print("Connected with result code " + str(rc))
    # 连接完成之后订阅gpio主题
    client.subscribe([('toggle', 2), ('setup', 2)])
 
# 消息推送回调函数
def on_message(client, userdata, msg):
    print("%s %s" % (msg.topic, msg.payload))
    # 获得负载中的pin 和 value
    data = json.loads(msg.payload)
    pin = data['pin']
    val = data['val']
    print( pin, val)
    if 'setup' == msg.topic:
        setup(str(pin), val)
    elif 'toggle' == msg.topic:
        if int(val) == 0:
            off(pin)
        else:
            on(pin)


if __name__ == '__main__':
    client = mqtt.Client()
    client.username_pw_set('admin', password='123123123')
    client.on_connect = on_connect
    client.on_message = on_message
    
    try:
        # 请根据实际情况改变MQTT代理服务器的IP地址
        client.connect("www.ruichen.top", 1883, 600)
        client.loop_forever()
    except KeyboardInterrupt:
        client.disconnect()
        if load_gpio:
            clean()