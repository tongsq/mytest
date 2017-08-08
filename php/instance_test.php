<?php
namespace Test;
interface UserInterface{
	public function setName(String $name);
	public function getName();
}
abstract class AbstractUser implements UserInterface{

	protected $name;
	const AGE = 18;
	public function getName(){
		echo 'name is ' . $this->name . "\n";
	}
	abstract public function setName(String $name);
}
class User extends AbstractUser{

	protected $name;
	public function setName(String $name){
		$this->name = $name;
	}
	//public function getName(){
	//	echo $this->name . "\n";
	//}
}
$user = new User();
$user->setName(123);
$user->getName();
var_dump($user instanceof UserInterface);
var_dump($user instanceof AbstractUser);
//var_dump(AGE);
var_dump($user::AGE);
