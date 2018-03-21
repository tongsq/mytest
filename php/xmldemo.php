<?php
$xml = '';
$obj = @simplexml_load_string($xml);
var_dump($obj);
foreach($obj as $name => $value) {
	echo $name . "\n";
	var_dump($value);
}
