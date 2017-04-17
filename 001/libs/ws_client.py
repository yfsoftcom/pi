import websocket
import thread
import time
import json
from common import get_ip

def on_message(ws, message):
  print 'message: ', message
  data = json.loads(message)
  print data
  if data['channel'] == 'ping':
    def run(*args):
      ip = get_ip()
      content = json.dumps({ 'id': 1, 'message': '', 'ip': ip, 'channel': 'upload'}, ensure_ascii = False)
      ws.send(content)

    thread.start_new_thread(run, ())

def on_error(ws, error):
  print error

def on_close(ws):
  print "### closed ###"

def on_open(ws):
  def run(*args):
    ws.send('{"channel":"ping"}')
    print "thread terminating..."
    
  thread.start_new_thread(run, ())


if __name__ == "__main__":
  websocket.enableTrace(True)
  ws = websocket.WebSocketApp("ws://10.1.129.107:10000",
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close)
  ws.on_open = on_open
  ws.run_forever()