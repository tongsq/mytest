<?php
$arr = array(1, 2, 'a' => 'a', 'b' => 'b');
var_dump(array_pop($arr));
var_dump($arr);
array_unshift($arr, 0);
var_dump($arr);
array_shift($arr);
var_dump($arr);
array_push($arr, 'c');
var_dump($arr);

$arr2 = array();
$arr2[8] = 8;
$arr2['a'] = 'a';
$arr2[1] = 1;
var_dump($arr2);
array_shift($arr2);
var_dump($arr2);
