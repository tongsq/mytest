<?php
$keys = array('foo', 5, 10, 'bar');
$a = array_fill_keys($keys, 'banana');
var_dump($a);
//array_fill
var_dump(array_fill(2, 3, 'banana'));
