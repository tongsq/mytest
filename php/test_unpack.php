<?php
$binarydata = "\x04\x00\xa0\x00";
var_dump($binarydata, strlen($binarydata));
$array = unpack("c4", $binarydata);
print_r($array);
$binarydata = pack("nvc*", 0x1234, 0x5678, 65, 66, 67);
var_dump($binarydata);
var_dump(chr(65));
