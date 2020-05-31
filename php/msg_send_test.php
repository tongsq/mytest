<?php
$queue_key = ftok(__FILE__, "a");
var_dump($queue_key);
$queue = msg_get_queue($queue_key);
$msg = 'aaa';
$msgtype = 1;
msg_send($queue, $msgtype);
$stat_arr = msg_stat_queue($queue);
var_dump($stat_arr);
var_dump(msg_receive($queue, 1, $msgtype, 1024, $data, true, MSG_IPC_NOWAIT));
var_dump($data);
