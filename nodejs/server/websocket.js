var http = require('http');
var io = require('socket.io');

var server = http.createServer((req, res)=>{
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.end('<h1>Hello Socket Lover!</h1>');
});
server.listen(8080);

var socket = io.listen(server);
socket.on('connection', function(client){
    client.on('message', function(event){
	console.log('Received message from client', event);
    });
    client.on('disconnect', function(){
	clearInterval(interval);
	console.log('Server has disconnected');
    });
    var interval = setInterval(function(){
	client.send('this is a message from the server!' + new Date().getTime());
    }, 5000);
});
