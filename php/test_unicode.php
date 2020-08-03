<?php
/**
 * utf8字符转换成Unicode字符
 * @param [type] $utf8_str Utf-8字符
 * @return [type]      Unicode字符
 */
function utf8_str_to_unicode($utf8_str) {
  $unicode = 0;
  $unicode = (ord($utf8_str[0]) & 0x1F) << 12;
  $unicode |= (ord($utf8_str[1]) & 0x3F) << 6;
  $unicode |= (ord($utf8_str[2]) & 0x3F);
  return dechex($unicode);
}
 
/**
 * Unicode字符转换成utf8字符
 * @param [type] $unicode_str Unicode字符
 * @return [type]       Utf-8字符
 */
function unicode_to_utf8($unicode_str) {
  $utf8_str = '';
  $code = intval(hexdec($unicode_str));
  //这里注意转换出来的code一定得是整形，这样才会正确的按位操作
  $ord_1 = 0xe0 | ($code >> 12);
  $ord_2 = 0x80 | (($code >> 6) & 0x3f);
  $ord_3 = 0x80 | ($code & 0x3f);
  $utf8_str = chr($ord_1) . chr($ord_2) . chr($ord_3);
  return $utf8_str;
}
$str = '你';
$str = utf8_str_to_unicode($str);
var_dump($str);
var_dump(unicode_to_utf8($str));
