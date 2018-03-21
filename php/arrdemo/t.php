<?php
$arr1 = range(1, 5);
$arr2 = range(1, 5);
shuffle($arr1);
shuffle($arr2);
var_dump($arr1, $arr2);
var_dump(array_merge($arr1, $arr2));
var_dump($arr1 + $arr2);
