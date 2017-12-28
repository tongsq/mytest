<?php
$serv = new swoole_server('0.0.0.0', 9500);
$serv->set(array('task_worker_num'=>4));
$serv->on('receive', function($serv, $fd, $from_id, $data){
	$task_id = $serv->task($data);
	echo "Dispath AsyncTask: id=$task_id\n";
	$serv->send($fd, 'ok');
});

$serv->on('task', function ($serv, $task_id, $from_id, $data){
	echo "new task[id=$task_id]\n";
	$serv->finish("{$data}->ok");
});
$serv->on('finish', function ($serv, $task_id, $data){
	echo "task finish:$data\n";
});
$serv->start();
