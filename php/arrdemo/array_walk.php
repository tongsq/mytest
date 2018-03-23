<?php
$arr = array('a' => 'aa', 'b' => 'bb', 'c' => 'cc');
function show(&$value, $key) {
	echo "key: {$key}, value: {$value}\n";
	$value = 'aa';
}
array_walk($arr, 'show');
var_dump($arr);
