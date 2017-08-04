<?php
//采用树结构过滤关键字
define('ENCODING', 'UTF-8');
$filter_words = array('fehe','fuck','fuce','fuc','你好');
$tree = array();
//生成过滤树
foreach($filter_words as $words){
    $node = createNode($words);
    mergeNode($node, $tree);
}
var_dump($tree);
$content = $_SERVER['argv'][1];
$len = mb_strlen($content, ENCODING);
for ($i=0; $i<$len; $i++)
{
    $j = $i;
    $tmp = $tree;
    while (true)
	{
		$current = mb_substr($content, $j, 1, ENCODING);
		if (isset($tmp[$current]))
		{
			if ($tmp[$current] === true){
				$content = mb_substr($content, 0, $i, ENCODING) . str_repeat('*',($j-$i+1)) . mb_substr($content, $j+1, NULL, ENCODING);
				$i = $j;
				break;
			}
			$tmp = $tmp[$current];
			$j++;
		}
		else{
			break;
		}
	}	
}
var_dump($content);
function createNode($words){
	if (mb_strlen($words, ENCODING) == 1){
		return array($words=>true);
	}
	else{
		return array(mb_substr($words, 0, 1, ENCODING)=>createNode(mb_substr($words, 1, NULL, ENCODING)));
	}
}
function mergeNode($node, &$tree){
	foreach($node as $k=>$v)
	{
		if ($v === true){
			$tree[$k] = true;
			break;
		}
		if (!isset($tree[$k])){
			$tree[$k] = $v;
			break;
		}
		mergeNode($node[$k], $tree[$k]);
		break;
	}
}
