var str = "456";
for(let s of str){
    console.log(s);
}
function str2num(str){
    var arr = [];
    for(let s of str){
	arr.push(s-0);
    }
    console.log(arr);
    return arr.reduce(function(x, y){
	return 10 * x + y; 	
    });
}
console.log(str2num("123456"));
