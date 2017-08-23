<?php
//pecl install xdebug
class Myclass{

	public function mycall($other){
		$other->call();
	}
}

class Otherclass{

	public function call(){
		echo "xdebug_call_class: " . xdebug_call_class();
		echo "\nxdebug_call_function: " . xdebug_call_function();
		echo "\nxdebug_call_file: " . xdebug_call_file();
		echo "\nxdebug_call_line: " . xdebug_call_line();
		echo "\n" . __CLASS__ . ":" . __FUNCTION__ . ":" . __LINE__ . "\n";
		var_dump(xdebug_get_function_stack());
	}
}
$myclass = new Myclass();
$other = new Otherclass();
$myclass->mycall($other);
$arr = array('a'=>1, 'b'=>2);
$arr = new ArrayIterator($arr);
var_dump($arr->current());
