<?php
$o = [];
@$var = ["",0,null,1,2,3,$foo,$o['myIndex']];
array_walk($var, function($v) {
    echo (!isset($v) || $v == false) ? 'true ' : 'false';
    echo ' ' . (empty($v) ? 'true' : 'false');
    echo "\n";
});
