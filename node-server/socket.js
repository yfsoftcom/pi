var io = require('socket.io')(10000);

io.on('connection', function (socket) {
  socket.emit('hi', { message: 'hi there world' });
  socket.on('sayHi', function (data) {
    console.log(data);
    io.emit('sayHi', data);
  });

  socket.on('disconnect', function () {
    io.emit('user disconnected');
  });
});
