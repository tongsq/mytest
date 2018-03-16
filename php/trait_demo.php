<?php
trait HelloWorld{
	public function say()
	{
		echo "trait say hello world\n";
		$this->say_private();
	}
	private function say_private()
	{
		echo "trait say_private hello world\n";
	}
}
class Base{
	public function say()
	{
		echo "Base say\n";	
	}
	private function sayBase(){
		echo "Base say private\n";
	}
}

class Demo extends Base
{
	use HelloWorld;
	public function sayDemo(){
		echo "demo say\n";
		$this->say_private();
		//parent::sayBase();
	}
}
$obj = new Demo();
$obj->say();
$obj->sayDemo();
