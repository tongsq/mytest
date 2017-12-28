<?php
$client = new swoole_client(SWOOLE_SOCK_TCP, SWOOLE_SOCK_ASYNC);
$client->on('connect', function($cli){
	$cli->send("hello");
});

$client->on('receive', function($cli, $data){
	echo "received: {$data}\n";
});
$client->on('error', function($cli){
	echo "connect close\n";
});
$client->on('close', function($cli){
	echo "connect close\n";
});

$client->connect('127.0.0.1', 9500, 0.5);
