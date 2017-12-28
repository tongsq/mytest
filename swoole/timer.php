<?php
$tick = swoole_timer_tick(1000, function ($timer_id){
	echo "tick-1000ms:{$timer_id}\n";
});
echo "{$tick} tick start\n";
$after = swoole_timer_after(3000, function(){
	global $after,$tick;
	echo "after 3000ms:{$after}\n";
	swoole_timer_clear($tick);
});
echo "after start\n";
