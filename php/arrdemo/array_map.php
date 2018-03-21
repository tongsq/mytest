<?php
$arr1 = array('a', 'b', 'c');
$arr2 = array('aa', 'bb', 'cc');
function demo($value1, $value2){
	return array($value1 => $value2);
}
$arr3 = array_map('demo', $arr1, $arr2);
var_dump($arr3);
