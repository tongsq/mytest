<?php
$arr = array('a' => 'aa', 'b' => 'bb', 'c' => 'cc');
$item = each($arr);
var_dump($item);
list($key, $val) = $item;
echo "key: {$key}, val: {$val}\n";
