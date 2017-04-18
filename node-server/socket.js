const WebSocket = require('ws');
 
const wss = new WebSocket.Server({ port: 10000 });

// Broadcast to all. 
wss.broadcast = function broadcast(data) {
  wss.clients.forEach(function each(client) {
    if (client.readyState === WebSocket.OPEN) {
      client.send(data);
    }
  });
};

wss.on('connection', function (ws) {
  console.log('client connected');
  ws.on('message', function (message) {
    const data = JSON.parse(message);
    console.log(data);
    if(data.channel=='login'){
      data.channel = 'online';
      wss.broadcast(JSON.stringify(data))
    }
  });
});
