<?php
function mod(int $x, int $n)
{
	$result = 1;
	for ($i=0; $i<$n; $i++) {
		$result *= $x;
	}
	return $result;
}
echo "2^10 = " . mod(2, 10) . "\n";
echo "2^16 = " . mod(2, 16) . "\n";
echo "2^32 = " . mod(2, 32) . "\n";
