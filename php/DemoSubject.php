<?php
class DemoSubject implements SplSubject {
	private $observers, $value;

	public function __construct(){
		$this->observers = new SplObjectStorage();
	}
	public function attach(SplObserver $observer){
		$this->observers->attach($observer);
	}
	public function detach(SplObserver $observer){
		$this->observers->detach($observer);
	}
	public function notify(){
		foreach($this->observers as $observer){
			$observer->update($this);
		}
	}
	public function setValue($value){
		$this->value = $value;
		$this->notify();
	}
	public function getValue(){
		return $this->value;
	}
}
class DemoObserver implements SplObserver {
	public function update (SplSubject $subject){
		echo "the new value is " . $subject->getValue()."\n";
	}
}
$demo = new DemoSubject();
$demo->attach(new DemoObserver);
$demo->setValue('hello');
