<?php
$ar1 = array(10, 100, 100, 0);
$ar2 = array(1, 3, 2, 4);
array_multisort($ar1, $ar2);

var_dump($ar1);
var_dump($ar2);

$arr3 = array(123, 456, 'abc', '92');
array_multisort($arr3, SORT_DESC, SORT_NUMERIC);
var_dump($arr3);
