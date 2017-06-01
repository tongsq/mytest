<?php
function create($start, $end){
    while($start <= $end){
	yield $start;
        $start++;
    }
}
$arr = create($argv[1], $argv[2]);
foreach($arr as $val){
    echo "{$val} ";
}
echo "\n";
