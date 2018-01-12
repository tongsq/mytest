<?php
//$serv = new Swoole\Http\Server('0.0.0.0', 9500);
$serv = new Swoole\Server('0.0.0.0', 9500);
$rk = new RdKafka\Producer();
$rk->setLogLevel(LOG_DEBUG);
$rk->addBrokers("10.10.16.129:9092,10.10.16.129:9093");
$topic = $rk->newTopic("test5");
$a = 1;
date_default_timezone_set('GMT');
$serv->on('workerstart', function($serv, $id) {
	$rk = new RdKafka\Producer();
    $rk->setLogLevel(LOG_DEBUG);
    $rk->addBrokers("10.10.16.129:9092,10.10.16.129:9093");
    $topic = $rk->newTopic("test5");
	$serv->topic = $topic;
	echo "init kafka client\n";
});
$serv->on('receive', function ($serv, $fd, $from_id, $data) {
    $serv->topic->produce(RD_KAFKA_PARTITION_UA, 0, "Message payload swoole 222");
	$data = "success\n";
	$length = strlen($data);
	$date = date('D, d M Y H:i:s e');
	$msg = <<<EOF
HTTP/1.1 200 OK
Server: swoole-http-server
Date: {$date}
Content-Type: text/html; charset=utf-8
Content-Length: {$length}
Connection: close

{$data}
EOF;
    $serv->send($fd, $msg);
    $serv->close($fd);
});
/*
$serv->on('request', function($request, $response){
	global $a;
	/*
	$rk = new RdKafka\Producer();
	$rk->setLogLevel(LOG_DEBUG);
	$rk->addBrokers("10.10.16.129:9092,10.10.16.129:9093");
	$topic = $rk->newTopic("test5");
	$re = $topic->produce(RD_KAFKA_PARTITION_UA, 0, "Message payload swoole 222");
	$response->header('Content-Type', 'text/html; charset=utf-8');
	$a++;
	$response->end("success {$a}");
});
*/
//$serv->set(['worker_num'=>1]);
$serv->start();
