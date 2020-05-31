<?php
class Parents{
	protected $name = '';
	
	function __construct($name){
		$this->name = $name;
	}
	public function say(){
		global $str;
		echo "my name is {$this->name}, str is {$str}\n";
	}
}

class Child extends Parents{
	public $name = '';
	
	function __construct($name, $age){
		$this->name = $name;
		$this->age = $age;
		Parent::__construct($name . " parent");
		$this->name = $name;
	}
	function __destruct(){
		echo "Child descruct\n";
	}
	function say2(){
		echo "my name is {$this->name}, age is {$this->age}\n";
	}
}
$str = 'abc';
$obj = new Child('tom', 18);
$obj->say();
$obj = null;
var_dump('ccccccccccc');
var_dump('bbbbbbbbbbb');
exit();
