<?php
$serv = new Swoole\Http\Server('0.0.0.0', 9500);
$atomic = new Swoole\Atomic(1);
$serv->on('request', function($request, $response) use ($atomic){
	//var_dump($request->get, $request->post);
	$response->header('Content-Type', 'text/html; charset=utf-8');
	$response->end($atomic->add(1) . "<h1>Hello Swoole.#" . rand(1000, 9999) . "</h1>");
	sleep(10000);
});
$serv->set(['worker_num'=>1]);
$serv->start();
