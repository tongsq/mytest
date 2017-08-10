<?php
class SingleClass{

	static $instance;
	private $name;
	private function __construct(String $name){
		$this->name = $name;
	}

	private function __clone(){}

	public static function getInstance($name){
		if (!(self::$instance instanceof self)){
			self::$instance = new self($name);
		}		
		return self::$instance;
	}
	
	public function getName(){
		echo "SingleClass echo {$this->name}\n";
	}
}
class OtherSingleClass extends SingleClass{
	private function __construct(String $name){
		$this->name = $name;
	}
	public function getName(){
		echo "OtherSingleClass echo {$this->name}\n";
	}
}

$class = OtherSingleClass::getInstance('tom');
var_dump($class);
$class->getName();
