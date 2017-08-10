<?php
class ExceptionTest extends Exception{

	public function __construct(String $msg='',int $code=0){
		parent::__construct($msg, $code);
	}
}
try{
	throw new ExceptionTest('test', -2);
}
catch(Exception $e){
	var_dump($e->getTrace());
	echo $e->__toString() . "\n";
}
