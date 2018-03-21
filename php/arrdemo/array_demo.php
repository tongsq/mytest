<?php
$arr1 = array(1,'b' => 2,'c' => 3);
//shuffle($arr1);
var_dump($arr1);
$arr2 = array(1,22,33);
var_dump(array_merge($arr1, $arr2));

$arr1 = $arr2 = range(1000, 2000);  
$arr_merged = array();  
//接下来随机给两个数组赋值  
shuffle($arr1);  
shuffle($arr2);  
  
$arr_merged = array_merge($arr1, $arr2);  
  
//var_dump($arr_merged); 
