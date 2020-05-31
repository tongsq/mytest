<?php
exec('wc -l /usr/local/openresty/nginx/logs/access.log', $out);
list ($line) = explode(' ', $out[0]);
echo "line is {$line}\n";
echo "ok\n";
