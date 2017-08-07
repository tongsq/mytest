<?php
$xml = '<?xml version="1.0" encoding="utf-8"?>
<unified_login>
<ret code="0" msg="登陆成功"/>
</unified_login>';
$dom = new DOMDocument();
$dom->loadXML($xml);
$ret = $dom->getElementsByTagName('ret');
echo 'code: ' . $ret[0]->getAttribute('code') . "\n";
echo 'msg: ' . $ret[0]->getAttribute('msg') . "\n";

//var_dump($dom->getAttributeNS('msg', 'ret'));
