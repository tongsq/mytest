<?php
header('X-Accel-Buffering: no');//立即输出
header('Access-Control-Allow-Origin:*');//js跨域问题
ignore_user_abort(true);
connection_status();//1,2,3
connection_aborted();
