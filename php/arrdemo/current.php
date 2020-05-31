<?php
$arr = array(0,1,2,3,4);
var_dump(current($arr));
var_dump(next($arr));
var_dump(current($arr));
foreach($arr as $val){
	echo $val;
}
