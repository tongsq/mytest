<?php
$serv = new swoole_websocket_server('0.0.0.0', 9500);
$serv->on('open', function ($serv, $request){
	var_dump($request->fd, $request->get, $request->server);
	$serv->push($request->fd, "hello welcome\n");
});
$serv->on('message', function ($serv, $frame){
	echo "Message: {$frame->data}\n";
	$serv->push($frame->fd, "server get {$frame->data}");
});
$serv->on('close', function ($serv, $fd){
	echo "client-{$fd} is closed\n";
});
$serv->start();
