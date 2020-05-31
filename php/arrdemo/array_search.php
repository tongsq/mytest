<?php
$arr = array('a' => 'a', 'b' => 'b', 'b', 'b', 'a');
$arr2 = array_search('a', $arr, true);
$arr3 = array_keys($arr, 'b');
var_dump($arr2, $arr3);
