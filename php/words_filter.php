<?php
//采用树结构过滤关键字
$filter_words = array('fehe','fuck','fuce','fuc','你好');
$tree = array();
//生成过滤树
foreach($filter_words as $words){
    $node = createNode($words);
    mergeNode($node, $tree);
}
var_dump($tree);
$content = $_SERVER['argv'][1];
$len = strlen($content);
for ($i=0; $i<$len; $i++)
{
    $j = $i;
    $tmp = $tree;
    while (true)
	{
		if (isset($tmp[$content[$j]]))
		{
			if ($tmp[$content[$j]] === true){
				for (; $i<=$j; $i++)
				{
					$content[$i] = '*';
				}
				$i = $j;
				break;
			}
			$tmp = $tmp[$content[$j]];
			$j++;
		}
		else{
			break;
		}
	}	
}
var_dump($content);
function createNode($words){
	if (strlen($words) == 1){
		return array($words=>true);
	}
	else{
		return array($words[0]=>createNode(substr($words, 1)));
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
