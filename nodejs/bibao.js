function lazy_sum(arr){
    var sum = function(){
	return arr.reduce(function(x, y){
	    return x + y;
	});
    }
    return sum;
}
var arr = [1,2,3];
var f = lazy_sum(arr);
console.log(f);
console.log(f());
console.log(f([4,5,6]));
