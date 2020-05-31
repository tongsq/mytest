<?php
$city  = "San Francisco";
$state = "CA";
$event = "SIGGRAPH";

$location_vars = array("city", "state");

$result = compact("event", "nothing_here", $location_vars);
print_r($result);

$arr = array('a'=>'aa', 'b'=>'bb');
$count = extract($arr);
echo "count: {$count}, a: $a, b: $b\n";
