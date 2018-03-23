<?php
$arr = array(1,2,3,4);
function reduce($carry, $item) {
	return $carry * $item;
}
$result = array_reduce($arr, 'reduce', 1);
var_dump($result);
