<?php
// Array representing a possible record set returned from a database
$arr = array(
    array(
        'id' => 2135,
        'first_name' => 'John',
        'last_name' => 'Doe',
    ),
    array(
        'id' => 3245,
        'first_name' => 'Sally',
        'last_name' => 'Smith',
    ),
    array(
        'id' => 5342,
        'first_name' => 'Jane',
        'last_name' => 'Jones',
    ),
    array(
        'id' => 5623,
        'first_name' => 'Peter',
        'last_name' => 'Doe',
    )
);
echo "array_column:\n";
var_dump(array_column($arr, 'first_name', 'id'));
$arr2 = array_combine(Array('a','a','b'), Array(1,2,3));
echo "array_combine:\n";
var_dump($arr2);
echo "array_count_values:\n";
$arr3 = array(1, "hello", 1, "world", "hello");
print_r(array_count_values($arr3));
