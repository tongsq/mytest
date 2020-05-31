<?php
$serv = new swoole_server('127.0.0.1', 9500, SWOOLE_PROCESS, SWOOLE_SOCK_UDP);
$serv->on('Packet', function($serv, $data, $clientInfo) {
	$serv->sendto($clientInfo['address'], $clientInfo['port'], "server ".$data);
	var_dump($clientInfo);
});
$serv->start();

