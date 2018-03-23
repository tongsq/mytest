<?php
$arr = array('a', 'b', 'c', 'd');
var_dump(array_slice($arr, 1, 2));
var_dump(array_slice($arr, 1, -2));
var_dump(array_slice($arr, -2, 1));

array_splice($arr, 2, 2, array(11,22));
var_dump($arr);
