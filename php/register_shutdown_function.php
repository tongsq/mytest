<?php
function shutdown($msg){
	echo "the process is shutdown with {$msg}\n";
}
register_shutdown_function('shutdown', 'hello');
echo "process start..\n";
$aa = new aa();
echo "aa\n";
die('bb');
