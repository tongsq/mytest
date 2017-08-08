<?php
class MyObject{

	public static function myMethod(){
		self::myOtherMethod();
		static::myOtherMethod();
	}

	static function myOtherMethod(){
		echo "called from MyObject\n";
	}
}

class MyOtherObject extends MyObject{

	static function myOtherMethod(){
		echo "called from MyOtherObject\n";
	}
}

MyOtherObject::myMethod();
