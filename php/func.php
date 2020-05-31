<?php
$str = "a\tb\tc\n";
$arr = explode("\t", trim($str));
$arr = ['1'=>'a','b'=>'b','c'=>'c', 'd'=>'d'];
list($a,$b,$c) = $arr;
var_dump($arr);
var_dump($a,$b,$c);
class Test{
	public $msg='aa';
}
$obj = new Test();
var_dump($obj->msg);
