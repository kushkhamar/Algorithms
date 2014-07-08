<?php 

function quick_sort($array)
{
	$length = count($array);
	
	if($length <= 1){
		return $array;
	}
	
	else{
		$pivot = $array[0];
		$left = $right = array();
		for($i = 1; $i < count($array); $i++)
		{
			if($array[$i] < $pivot){
				$left[] = $array[$i];
			}
			else{
				$right[] = $array[$i];
			}
		}
		return array_merge(quick_sort($left), array($pivot), quick_sort($right));
	}
}

// here
$array = array(9,8,7,6,5,4,3,2,1,0,10,1000,0);
$sorted = quick_sort($array);
print_r($sorted);

?>
 
