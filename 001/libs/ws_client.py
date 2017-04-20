#!/usr/bin/env python
#encoding: utf-8  
import thread
import time
import json
from common import get_ip

from socketIO_client import SocketIO, BaseNamespace

class Namespace(BaseNamespace):

  def on_connect(self):
    self.emit('login',{ 'ip': get_ip(), 'id': 1, 'channel': 'login'})

  def on_message(self, message):
    print('on_message', message)

  def on_reconnect(self):
    print('[Reconnected]')

  def on_disconnect(self):
    print('[Disconnected]')


if __name__ == "__main__":
  socketIO = SocketIO('api.yunplus.io', 80, Namespace)
  socketIO.wait()