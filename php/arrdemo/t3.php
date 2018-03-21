<?php
$arr1 = array(100 => 'a', 200 => 'b');
$arr2 = array(5=>'c', 8=>'d');
var_dump($arr1, $arr2);
var_dump(array_merge($arr1, $arr2));
var_dump(array_merge($arr2, $arr1));
