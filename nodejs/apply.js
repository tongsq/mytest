function add(x, y){
    return parseInt(x) + parseInt(y);
}
var addTmp = add;
var count = 0;
var add = function(){
    count++;
    return addTmp.apply(null, arguments);
    return addTmp.call(null, x, y);
}
console.log(add(1, 2));
console.log(add(3, 4));
console.log("the count is: "+count);
