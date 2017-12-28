<?php
$client = new swoole_client(SWOOLE_SOCK_TCP);

if (!$client->connect('127.0.0.1', 9500, 0.5)){
	die("connect failed.\n");
}
if (!$client->send("hello")){
	die("send failed.\n");
}
$data = $client->recv();
if (!$data){
	die("recv failed.\n");
}
echo $data;
$client->close();
