function sum(...rest){
    var sum = 0;
    rest.forEach(function(value){
	sum += value;
    });
    return sum;
}
//console.log(sum(1, 2, 3));
function define(){
    a = 100;
}
let a = 10;
console.log(a);
define();
console.log(a);
