const WebSocket = require('ws');
socket = new WebSocket("ws://10.10.16.129:8080");
			socket.onopen = (event)=>{
			    socket.send("i'am the client and i'am listening");
				socket.onmessage = (event)=>{
					console.log("client received a message", event);
				}
				socket.onclose = (event)=>{
					console.log("client notified socket has closed", event);
				}
			};
