<?php
$serv = new swoole_server('127.0.0.1', 9500);
$serv->on('connect', function ($serv, $fd){
	echo "client: connect.\n";
});

$serv->on('receive', function ($serv, $fd, $from_id, $data) {
	var_dump($fd, $from_id, $data);
	$serv->send($fd, "server: " . $data);
	if ($data == "exit\r\n"){
		$serv->close($fd);
	}
});

$serv->on('close', function ($serv, $fd) {
	echo "Client: close.\n";
});
$serv->start();
