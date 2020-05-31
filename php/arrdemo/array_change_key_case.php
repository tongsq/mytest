<?php
$arr = array('aBc' => 'abc', 'Def' => 'def');
print_r(array_change_key_case($arr, CASE_UPPER));
print_r(array_change_key_case($arr, CASE_LOWER));
