<?php
namespace Test;
$o = [];
@$var = ["",0,null,1,2,3,$foo,$o['myIndex']];
array_walk($var, function($v) {
    echo (!isset($v) || $v == false) ? 'true ' : 'false';
    echo ' ' . (empty($v) ? 'true' : 'false');
    echo "\n";
});
if (true == -1){
	echo 'true == -1' . "\n";
}
if (true == 1){
	echo 'true == 1' . "\n";
}
if (true != 0){
	echo 'true != 0' . "\n";
}
var_dump(__DIR__);
var_dump(dirname(__DIR__));
var_dump(dirname(__FILE__));
var_dump(__FILE__);
var_dump(__NAMESPACE__ . '\\');
