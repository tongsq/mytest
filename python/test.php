<?php
function fact(float $n){
	if ($n == 1)
		return 1;
	return fact($n-1) * $n;
}
function newfact(int $n, float $product=1){
	if ($n == 1)
		return $product;
	return newfact($n-1, $n * $product);
}
//print("fact(1000): " . fact(1000) . "\n");
print("newfact(100): " . newfact(100) . "\n");
$f = 'fact';
print($f(100));
$arr = [1,2,3,4,5];
function func($num){
	return $num + 1;
}
$arr2 = array_map('func', $arr);
var_dump($arr2);
