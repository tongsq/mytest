'use strict';
const WebSocket = require('ws');
const WebSocketServer = WebSocket.Server;
const wss = new WebSocketServer({
    port: 8080
});
wss.on('connection', function(ws){
    console.log(`[SERVER] connection()`);
    ws.on('message', function(message){
		console.log(`[SERVER] Received: ${message}`);
		ws.send(`server accept ${message}`, (err)=>{
	    	if (err){
			console.log(`[SERVER] error: ${err}`);
	    	}
		})
    });
    var interval = setInterval(()=>{
	ws.send('this is a message send from server' + new Date().getTime(), (err)=>{
	    if (err)
	        console.log(`send error: ${err}`);
	});
    }, 5000);
    ws.on('close', ()=>{
	clearInterval(interval);
	console.log('lost connect');
	console.log(ws);
    });
})
